// Object representing students and their grades
const students = {
    John: [85, 90, 78]
};

// 1. Display each student's name using for...in loop
console.log("Student Names:");
for (let student in students) {
    console.log(student);
}

console.log("\nStudent Grades:");
// 2. Display each student's grades using for...of loop
for (let student in students) {
    let grades = students[student];
    console.log(`${student}'s Grades:`);
    for (let grade of grades) {
        console.log(grade);
    }
}

console.log("\nAverage Grades:");
// 3. Calculate and display the average grade for each student
for (let student in students) {
    let grades = students[student];
    let sum = 0;
    for (let grade of grades) {
        sum += grade;
    }
    let avg = sum / grades.length;
    console.log(`${student}'s Average Grade: ${avg.toFixed(2)}`);
}
