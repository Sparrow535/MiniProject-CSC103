function validateForm() {
    var fname = document.getElementById('fname').value;
    var lname = document.getElementById('lname').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    var confirmPassword = document.getElementById('confirm-password').value;
    var genderMale = document.getElementById('male').checked;
    var genderFemale = document.getElementById('female').checked;

    document.getElementById('fname-error').innerHTML = '';
    document.getElementById('lname-error').innerHTML = '';
    document.getElementById('email-error').innerHTML = '';
    document.getElementById('password-error').innerHTML = '';
    document.getElementById('confirm-password-error').innerHTML = '';
    document.getElementById('gender-error').innerHTML = '';

    var isValid = true;

    if (fname.trim() === '') {
        document.getElementById('fname-error').innerHTML = 'First name is required';
        isValid = false;
    }

    if (lname.trim() === '') {
        document.getElementById('lname-error').innerHTML = 'Last name is required';
        isValid = false;
    }

    if (email.trim() === '') {
        document.getElementById('email-error').innerHTML = 'Email is required';
        isValid = false;
    } else if (!isValidEmail(email)) {
        document.getElementById('email-error').innerHTML = 'Invalid email address';
        isValid = false;
    }

    if (password.trim() === '') {
        document.getElementById('password-error').innerHTML = 'Password is required';
        isValid = false;
    }

    if (confirmPassword.trim() === '') {
        document.getElementById('confirm-password-error').innerHTML = 'Please confirm your password';
        isValid = false;
    } else if (password !== confirmPassword) {
        document.getElementById('confirm-password-error').innerHTML = 'Passwords do not match';
        isValid = false;
    }

    if (!genderMale && !genderFemale) {
        document.getElementById('gender-error').innerHTML = 'Please choose gender';
        isValid = false;
    }

    return isValid;
}

function isValidEmail(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
}

function signUp(event) {
    event.preventDefault(); // Prevent the default form submission

    if (!validateForm()) {
        return;
    }

    var gender = document.getElementById('male').checked ? 'Male' : 'Female';

    var _data = {
        first_name: document.getElementById('fname').value,
        last_name: document.getElementById('lname').value,
        email: document.getElementById('email').value,
        gender: gender,
        password: document.getElementById('password').value
    };

    fetch('/signup', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    .then(response => {
        if (response.status == 201) {
            window.open("index.html", "_self");
        } else {
            return response.json().then(data => {
                throw new Error(data.message);
            });
        }
    })
    .catch(error => {
        alert(error);
    });
}

document.getElementById('signup').addEventListener('submit', signUp);
s