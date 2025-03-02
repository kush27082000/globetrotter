//******** */
// Function to get cookies as an object
// Function to get cookies as an object
function getCookies() {
    console.log("document.getCookies starts")
    console.log(document.cookie.split(';'))
    console.log("document.getCookies ends")
    return document.cookie.split('; ').reduce((acc, cookie) => {
        const [name, value] = cookie.split('=');
        acc[name] = value;
        console.log("acc")
        if (name === "username") {
            console.log("Username cookie found:", acc[name]);
            // map[username] = acc[name]
        }    
        return acc
    }, {});
}

function showGameUI() {
    document.getElementById("game").style.display = "block";
    document.getElementById("auth").style.display = "none";
}

function showAuthUI() {
    document.getElementById("game").style.display = "none";
    document.getElementById("auth").style.display = "block";
}

function showError(message) {
    alert(message);
}

function register() {
    const Username = document.getElementById("username").value.trim();
    if (!Username) {
        showError("Username cannot be empty!");
        return;
    }
    
    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ Username })
    })
    .then(res => res.json())
    .then(data => {
        alert(data.message);
        document.cookie = `username=${Username}; path=/;`;
        showGameUI();
        getScore();
    })
    .catch(error => showError(error.message));
}

function login() {
    const username = document.getElementById("username").value.trim();
    if (!username) {
        showError("Username cannot be empty!");
        return;
    }
    
    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username })
    })
    .then(res => {
        if (!res.ok) {
            throw new Error("Login failed - User does not exist. Kindly register!");
        }
        return res.json(); // ✅ Ensure JSON is properly returned
    })
    .then(data => {
        if (!data || !data.message) {
            throw new Error("Unexpected response from server!");
        }
        
        alert(data.message);
        document.cookie = `username=${username}; path=/;`; // ✅ Fixed syntax
        showGameUI();
        getScore();
    })
    .catch(error => showError(error.message || "An error occurred while logging in."));
}


function fetchAPI(endpoint, method = "GET", body = null) {
    const cookies = getCookies();
    const username = cookies["username"] || "";

    console.log("Cookie: username=" + username);

    const options = {
        method,
        headers: {
            "Content-Type": "application/json"
            // "Cookie": `username=${username}`,
            //  "credentials": "include"
        },
    };
    console.log(options)

    if (body) options.body = JSON.stringify(body);

    return fetch(`http://localhost:8080/${endpoint}?username=${username}`, options)
        .then(res => res.json());
}

let activeClueRequests = new Map(); // Store clue requests (username -> clueId mapping)

function getClue() {
    const cookies = getCookies();
    const username = cookies["username"];
    if (!username) {
        showError("User not authenticated!");
        return;
    }

    // Fetch the clue and store the request
    const clueRequest = fetchAPI("clue").then(data => {
        if (data.clue) {
            activeClueRequests.set(username, data.clue.ID); // Store clueId mapped to user
            document.getElementById("clue-text").innerText = data.clue.clues || "No clue available!";
        } else {
            showError("No clue available!");
        }
    }).catch(error => showError(error.message));

    return clueRequest; // Return the promise to allow chaining
}

function submitAnswer() {
    const answer = document.getElementById("answer").value.trim();
    if (!answer) {
        alert("Please enter an answer!");
        return;
    }

    const cookies = getCookies();
    const username = cookies["username"];
    const clue_id = activeClueRequests.get(username); // Get clue ID specific to user

    if (!clue_id) {
        showError("No active clue! Please get a new clue first.");
        return;
    }

    // Create an array of promises for concurrent execution
    const answerRequest = fetchAPI("submit", "POST", { username, answer, clue_id });

    const scoreRequest = getScore(); // Fetch the latest score concurrently

    Promise.all([answerRequest, scoreRequest])
        .then(([answerData, scoreData]) => {
            document.getElementById("feedback").innerText = answerData.result === "correct" ? 
                `✅ Correct! Fun Fact: ${answerData.fun_facts || "No fact available."} | Trivia: ${answerData.trivia || "No trivia available."}` : 
                `❌ Incorrect! Try again!`;
        })
        .catch(error => showError(error.message));
     
    getScore()     
}

function getScore() {
    fetchAPI("score").then(data => {
        document.getElementById("score").innerText = data.score || 0;
    }).catch(error => showError(error.message));
}

function resetScore() {
    fetchAPI("reset", "POST").then(() => {
        alert("Score reset!");
        getScore();
    }).catch(error => showError(error.message));
}

function inviteFriend() {
    const cookies = getCookies();
    const username = cookies["username"];
    fetchAPI("invite", "POST").then(data => {
        // const inviteLink = `http://localhost:8080/invite/${data.inviteCode}`;
        const inviteLink = data.invite_link
        document.getElementById("invite-link").innerHTML = `Invite Link: <a href="${inviteLink}" target="_blank">${inviteLink}</a>`;
    }).catch(error => showError(error.message));
}

function logout() {
    document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    alert("Logged out successfully!");
    location.reload();
}

document.addEventListener("DOMContentLoaded", () => {
    const cookies = getCookies();
    // if (cookies.username) {
    //     showGameUI();
    //     getScore();
    // } else {
    //     showAuthUI();
    // }
    showAuthUI();
});
