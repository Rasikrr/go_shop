const forms = document.querySelector(".forms"),
    pwShowHide = document.querySelectorAll(".eye-icon"),
    links = document.querySelectorAll(".link");

pwShowHide.forEach(eyeIcon => {
    eyeIcon.addEventListener("click", () => {
        let pwFields = eyeIcon.parentElement.parentElement.querySelectorAll(".password");

        pwFields.forEach(password => {
            if(password.type === "password"){
                password.type = "text";
                eyeIcon.classList.replace("bx-hide", "bx-show");
                return;
            }
            password.type = "password";
            eyeIcon.classList.replace("bx-show", "bx-hide");
        })

    })
})

links.forEach(link => {
    link.addEventListener("click", e => {
        e.preventDefault(); //preventing form submit
        forms.classList.toggle("show-signup");
    })
})

// Handle form submissions
const loginForm = document.querySelector('.form.login form');
const signupForm = document.querySelector('.form.signup form');

loginForm.addEventListener('submit', function(e) {
    e.preventDefault();
    const email = this['email-signin'].value;
    const password = this['password-signin'].value;

    fetch('http://localhost:8080/auth/signin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    }).then(response => response.json())
        .then(data => {
            console.log(data);
            if(data && data.error) {
                document.getElementById('login-error').textContent = "Email or password is invalid";
            } else {
                window.location.href = "/";
            }
        }).catch(error => console.error('Error:', error));
});
signupForm.addEventListener('submit', function(e) {
    e.preventDefault();
    const name = this['name-signup'].value;
    const email = this['email-signup'].value;
    const password1 = this['password1-signup'].value;
    const password2 = this['password2-signup'].value;

    fetch('http://localhost:8080/auth/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({name, email, password1, password2})
    }).then(response => response.json())
        .then(data => {
            console.log("HERE", data);
            if (data && data.error) {
                console.log(data.error.message)
                document.getElementById('signup-error').textContent = data.error;
            } else {
                window.location.href = "/";
            }
        })
});