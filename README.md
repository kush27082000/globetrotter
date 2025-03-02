# globetrotter
//
show databases;



CREATE DATABASE IF NOT EXISTS globetrotter;
USE globetrotter;

-- Table: destinations
CREATE TABLE destinations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    clues JSON NOT NULL,
    fun_facts JSON NOT NULL,
    trivia JSON NOT NULL
);

-- Table: users
-- CREATE TABLE users (
    -- id INT AUTO_INCREMENT PRIMARY KEY,
    -- username VARCHAR(100) UNIQUE NOT NULL,
    -- score INT DEFAULT 0,
    -- created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	score INT DEFAULT 0
);

INSERT INTO destinations (city, country, clues, fun_facts, trivia) VALUES
('London', 'UK',
  '["Home to a famous clock tower often mistaken for its bell.",
    "A city where the Queen used to reside in Buckingham Palace."]',
  '["Big Ben actually refers to the bell, not the tower.",
    "London has over 170 museums, including the British Museum."]',
  '["The London Underground is the oldest metro system in the world.",
    "More than 300 languages are spoken in London."]'),

('Rome', 'Italy',
  '["A city with a famous amphitheater once used for gladiator fights.",
    "Known as the \'Eternal City\' and home to an independent country inside it."]',
  '["The Colosseum could hold up to 50,000 spectators!",
    "Vatican City, the smallest country in the world, is inside Rome."]',
  '["Romans invented concrete, revolutionizing architecture.",
    "Rome has more fountains than any other city in the world."]'),

('Berlin', 'Germany',
  '["This city was divided by a wall for nearly 30 years.",
    "Home to the Brandenburg Gate and a vibrant art scene."]',
  '["The Berlin Wall fell in 1989, marking German reunification.",
    "Berlin has more bridges than Venice!"]',
  '["The city is home to the world’s largest beer garden.",
    "Berlin’s Museum Island is a UNESCO World Heritage site."]'),

('Rio de Janeiro', 'Brazil',
  '["Famous for its giant Christ the Redeemer statue.",
    "Host of one of the world’s most famous carnival festivals."]',
  '["Sugarloaf Mountain offers a stunning view of the city.",
    "Rio’s beaches, like Copacabana and Ipanema, are world-famous."]',
  '["The city was once the capital of Portugal.",
    "Rio’s Maracanã Stadium hosted the 1950 and 2014 World Cup finals."]'),

('Cape Town', 'South Africa',
  '["This city has a famous flat-topped mountain.",
    "Located at the meeting point of the Atlantic and Indian Oceans."]',
  '["Table Mountain is over 260 million years old.",
    "Cape Town was the first city in the world to have a botanical garden."]',
  '["Penguins live on the beaches near the city.",
    "Cape Town was the first non-European city to get a Nobel Square."]');



//
backend-
go run main.go

frontend
run index.html


//git - 
1. 
ERROR: Permission to kush27082000/globetrotter.git denied to deploy key
fatal: Could not read from remote repository.
Please make sure you have the correct access rights
    (or)
remote: Permission to kush27082000/globetrotter.git denied to ksoniAngel.
fatal: unable to access 'https://github.com/kush27082000/globetrotter.git/': The requested URL returned error: 403

//git remote set-url origin  https://PAT@github.com/username/reponame.git
after that can do git push
 

2. 
 kushagra.soni@24F-MACABL-1244 globaltrotter % git push --set-upstream origin main
To https://github.com/kush27082000/globetrotter.git
 ! [rejected]        main -> main (fetch first)
error: failed to push some refs to 'https://github.com/kush27082000/globetrotter.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

do->  git pull --rebase origin main
and after that can do git push