var userContainer = document.getElementById('user')

if (userContainer != null) {
    var navMenuContainer = document.getElementById('nav-menu-container')

    navMenuContainer.addEventListener('mouseenter', function() {
        document.getElementById('nav-menu-container').style.display = 'block';
    })

    navMenuContainer.addEventListener('mouseleave', function() {
        document.getElementById('nav-menu-container').style.display = 'none';
    })

    userContainer.addEventListener('mouseenter', function() {
        document.getElementById('nav-menu-container').style.display = 'block';
    })

    userContainer.addEventListener('mouseleave', function() {
        document.getElementById('nav-menu-container').style.display = 'none';
    })
}