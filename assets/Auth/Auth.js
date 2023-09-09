const switchers = [...document.querySelectorAll('.switcher')]
const loginBtn = document.querySelector('#btn-login')
const signupBtn = document.querySelector('#btn-signup')
const loginForm = document.querySelector('#login-form')
const signupForm = document.querySelector('#signup-form')
const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/
const usernameRegex = /^[a-z0-9A-Z]{3,16}$/
const isNonWhiteSpace = /^\S*$/;
const isContainsUppercase = /^(?=.*[A-Z]).*$/;
const isContainsLowercase = /^(?=.*[a-z]).*$/;
const isContainsNumber = /^(?=.*[0-9]).*$/;
const isContainsSymbol = /^(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]).*$/;
const isValidLength = /^.{6,16}$/;
// const passRegexLower = /.*[a-z].*/
// const passRegexUpper = /.*[A-Z].*/
// const passRegexLength = /.*[a-zA-Z0-9]{8,}.*/
let regularExpression = /^(\S)(?=.*[0-9])(?=.*[A-Z])(?=.*[a-z])(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹])[a-zA-Z0-9~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]{10,16}$/;
switchers.forEach(item => {
    item.addEventListener('click', function() {
        switchers.forEach(item => item.parentElement.classList.remove('is-active'))
        this.parentElement.classList.add('is-active')
    })
})

function Switcher() {
    switchers.forEach(item => switchers[1].parentElement.classList.remove('is-active'))
    switchers[0].parentElement.classList.add('is-active')
}

loginBtn.onclick = (e => {
    e.preventDefault()
    let check = true;
    const formData = new FormData(loginForm);
    const email = formData.get('email');
    const password = formData.get('password');

    if(email.length == 0){
        document.getElementById('login-email-err').innerHTML = 'Please enter email';
        check = false;
    } else if (!emailRegex.test(email)) {
        document.getElementById('login-email-err').innerHTML = 'Invalid email';
        check = false;
    }else{
        document.getElementById('login-email-err').innerHTML = '';
    }
    if(password.length <6){
        document.getElementById('login-password-err').innerHTML = 'Password must be at least 6 characters';
        check = false;
    } else {
        document.getElementById('login-password-err').innerHTML = '';
    }

    if(check == false){
        return false;
    }
    $.ajax({
        url: '/api/users/login',
        method: 'POST',
        data: {
            email: email,
            password: password
        },
        success: function(data) {
            window.location.href = "/"
        },
        error: function(err) {
            if(err.status == 400){
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.responseJSON.message,
                })
            }
        }
    })
})

signupBtn.onclick = (e =>{
    e.preventDefault()
    let check = true;
    const formData = new FormData(signupForm);
    const username = formData.get('user_name');
    const email = formData.get('email');
    const password = formData.get('password');
    const cpassword = formData.get('cpassword');


    if (username.length == 0){
        document.getElementById('signup-username-err').innerHTML = 'Please enter username';
        check = false;
    }else if (!usernameRegex.test(username)){
        document.getElementById('signup-username-err').innerHTML = 'Invalid username';
        check = false;
    }else{
        document.getElementById('signup-username-err').innerHTML = '';
    }

    if (email.length === 0){
        document.getElementById('signup-email-err').innerHTML = 'Please enter email';
        check = false;
    } else if (!emailRegex.test(email)){
        document.getElementById('signup-email-err').innerHTML = 'Invalid email';
        check = false;
    }else {
        document.getElementById('signup-email-err').innerHTML = '';
    }


    if(!isNonWhiteSpace.test(password)){
        document.getElementById('signup-password-err').innerHTML = 'Password must not contain Whitespaces.';
        check = false;
    }
    if (!isContainsUppercase.test(password)) {
        document.getElementById('signup-password-err').innerHTML = "Password must have at least one Uppercase Character.";
        check = false;
    }
    if (!isContainsLowercase.test(password)) {
        document.getElementById('signup-password-err').innerHTML = "Password must have at least one Lowercase Character.";
        check = false;
    }
    if (!isContainsNumber.test(password)) {
        document.getElementById('signup-password-err').innerHTML = "Password must contain at least one Digit.";
        check = false;
    }
    if (!isContainsSymbol.test(password)) {
        document.getElementById('signup-password-err').innerHTML = "Password must contain at least one Special Symbol.";
        check = false;
    }
    if (!isValidLength.test(password)) {
        document.getElementById('signup-password-err').innerHTML = "Password must be 6-16 Characters Long.";
        check = false;
    }

    if(regularExpression.test(password)) {
        document.getElementById('signup-password-err').innerHTML ='';
    }



    if (cpassword != password){
        document.getElementById('signup-cpassword-err').innerHTML = 'Passwords do not match';
        check = false;
    }else {
        document.getElementById('signup-cpassword-err').innerHTML = '';
    }

    if(check == false){
        return false;
    }

    $.ajax({
        url: '/api/users/register',
        method: 'POST',
        data: {
            user_name: username,
            email: email,
            password: password,
            cpassword: cpassword
        },success: function(data) {
            console.log(data);
            Switcher();
        },error: function(err) {
            if(err.status == 400){
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.responseJSON.message,
                })
            }
        }

    })
})