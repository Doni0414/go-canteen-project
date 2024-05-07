formLoginBtn = document.getElementById('form-login-btn')
formSignupBtn = document.getElementById('form-signup-btn')

function toggleLoginSignup(toggleTo) {
    var loginFields = document.querySelectorAll('#login-inputs .field')
    var signupFields = document.querySelectorAll('#reg-inputs .field')
     
    if (toggleTo === 'signup') {
        document.getElementById('form-login-btn').classList.remove('form-clicked-btn')
        document.getElementById('form-signup-btn').classList.add('form-clicked-btn')
        document.getElementById('reg-form').action = '/auth?signup'
        document.getElementById('reg-inputs').style.display = 'block'
        document.getElementById('login-inputs').style.display = 'none'
        document.getElementById('registration-btn').textContent = 'Регистрация'
        document.getElementById('auth-type').value = 'signup'

        loginFields.forEach(field => {
            field.removeAttribute('required')
        })

        signupFields.forEach(field => {
            field.setAttribute('required', 'true')
        })
    } else if (toggleTo === 'login'){
        document.getElementById('form-login-btn').classList.add('form-clicked-btn')
        document.getElementById('form-signup-btn').classList.remove('form-clicked-btn')
        document.getElementById('reg-form').action = '/auth?login'
        document.getElementById('reg-inputs').style.display = 'none'
        document.getElementById('login-inputs').style.display = 'block'
        document.getElementById('registration-btn').textContent = 'Логин'
        document.getElementById('auth-type').value = 'login'

        signupFields.forEach(field => {
            field.removeAttribute('required')
        })

        loginFields.forEach(field => {
            field.setAttribute('required', 'true')
        })
    }
}

if (formLoginBtn != null) {
    formLoginBtn.addEventListener('click', function (event) {
        event.preventDefault()
        if (!formLoginBtn.classList.contains('form-clicked-btn')) {
            toggleLoginSignup('login')
        }
    })
}

if (formSignupBtn != null) {
    formSignupBtn.addEventListener('click', function (event) {
        event.preventDefault()
        if (!formSignupBtn.classList.contains('form-clicked-btn')) {
            toggleLoginSignup('signup')
        }
    })
}

const urlParams = new URLSearchParams(window.location.search)

if (urlParams.has("login")) {
    toggleLoginSignup('login')
} else {
    if (formSignupBtn != null) {
        toggleLoginSignup('signup')
    }
}

var form = document.getElementById('reg-form')
if (form != null) {
    var fields = form.getElementsByClassName('field')
    Array.from(fields).forEach(field => {
        field.addEventListener("keypress", function(event) {
            if (event.key === "Enter" || event.keyCode === 13) {
                form.submit()
            }
        })
    })
}