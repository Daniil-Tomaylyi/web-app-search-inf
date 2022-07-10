let startButton = document.getElementById('start-btn')
let menuList=document.querySelector('.navbar-nav')
let menuBtn = document.querySelector('.bi-menu-down')

startButton.onclick = function () {
    window.scrollTo({
        top: window.innerHeight,
        behavior: 'smooth'
    })
}
