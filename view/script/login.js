function validateForm() {
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
  
    document.getElementById('email-error').innerHTML = '';
    document.getElementById('password-error').innerHTML = '';
  
    var isValid = true;
  
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
  
    return isValid;
  }
  
  function isValidEmail(email) {
    var re = /\S+@\S+\.\S+/;
    return re.test(email);
  }
  
  function login() {
    if (!validateForm()) {
      return;
    }
  
    var _data = {
      email: document.getElementById("email").value,
      password: document.getElementById("password").value
    }
  
    fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=UTF-8'
      },
      body: JSON.stringify(_data)
    })
      .then(response => {
        if (!response.ok) {
          throw new Error(response.statusText);
        }
        return response.json(); // Ensure the response is parsed correctly
      })
      .then(data => {
        console.log('Response Data:', data); // Debug log
        if (data.gender === 'Male') {
          window.open("adminBoy.html", "_self");
        } else if (data.gender === 'Female') {
          window.open("adminGirl.html", "_self");
        } else {
          throw new Error('Unknown gender');
        }
      })
      .catch(error => {
        console.error('Login Error:', error); // Debug log
        if (error.message === "Unauthorized") {
          alert(error + ". Credentials do not match!");
          return
        }
      });
  }
  
  function logout() {
    fetch('/logout', {
      method: 'POST'
    })
      .then(response => {
        if (response.ok) {
          window.open("index.html", "_self");
        } else {
          throw new Error(response.statusText);
        }
      })
      .catch(error => {
        alert('Logout Error: ' + error.message);
      });
  }
  
  