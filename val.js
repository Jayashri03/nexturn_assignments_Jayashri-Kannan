function validateForm() {
    const name = document.getElementById('name').value.trim();
    const email = document.getElementById('email').value.trim();
    const age = document.getElementById('age').value.trim();
    let isValid = true;

    if (name.length < 3) {
        document.getElementById('nameError').innerText = 'Name must be at least 3 characters long';
        isValid = false;
    } else {
        document.getElementById('nameError').innerText = '';
    }

    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailPattern.test(email)) {
        alert('Please enter a valid email address');
        isValid = false;
    }

    if (age < 18 || age > 100 || isNaN(age)) {
        alert('Age must be a number between 18 and 100');
        isValid = false;
    }

    return isValid;
}