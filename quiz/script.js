let questions = [
    {
        question: "What is the largest planet in our solar system?",
        options: ["Earth", "Saturn", "Jupiter", "Uranus"],
        correct: 2
    },
    {
        question: "What is the chemical symbol for water?",
        options: ["O2", "H2O", "CO2", "H2"],
        correct: 1
    },
    {
        question: "Which animal is known as the King of the Jungle?",
        options: ["Tiger", "Elephant", "Lion", "Giraffe"],
        correct: 2
    },
    {
        question: "Which is the smallest prime number?",
        options: ["0", "1", "2", "3"],
        correct: 2
    },
    {
        question: "Who wrote 'Hamlet'?",
        options: ["Charles Dickens", "J.K. Rowling", "William Shakespeare", "Mark Twain"],
        correct: 2
    },
    {
        question: "What is the boiling point of water in Celsius?",
        options: ["100°C", "0°C", "50°C", "200°C"],
        correct: 0
    }
];

let currentQuestion = 0;
let score = 0;

document.getElementById("quiz-form").addEventListener("submit", submitForm);

function submitForm(event) {
    event.preventDefault();
    let selectedOption = document.querySelector("input[name='option']:checked");

    if (selectedOption) {
        let correctOption = questions[currentQuestion].correct;
     
        if (parseInt(selectedOption.value) === correctOption) {
            score++;
        }
        currentQuestion++;
        if (currentQuestion < questions.length) {
            displayQuestion();
        } else {
            displayResult();
        }
    } else {
        alert("Please select an option.");
    }
}

function displayQuestion() {
    let questionText = document.getElementById("question-text");
    let optionsList = document.getElementById("options");
   
    questionText.textContent = questions[currentQuestion].question;

    optionsList.innerHTML = "";

    questions[currentQuestion].options.forEach((option, index) => {
        let optionElement = document.createElement("li");
        optionElement.innerHTML = `
            <input type="radio" id="option${index + 1}" name="option" value="${index}">
            <label for="option${index + 1}">${option}</label>
        `;
        optionsList.appendChild(optionElement);
    });
}

function displayResult() {
    let resultElement = document.getElementById("result");
    resultElement.textContent = `You scored ${score} out of ${questions.length}!`;
}


displayQuestion();
