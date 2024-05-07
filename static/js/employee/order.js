var nextBtn = document.getElementById('next-btn')
var rejectBtn = document.getElementById('reject-btn')

document.getElementById('exit-icon').addEventListener('click', () => {
    var commentContainer = document.getElementById('comment-container')
    commentContainer.style.opacity = 0
    commentContainer.style.visibility = 'hidden'
})

function clickBtn(value) {
    var commentContainer = document.getElementById('comment-container')
    commentContainer.style.opacity = 1
    commentContainer.style.visibility = 'visible'

    if (value === "next") {
        commentContainer.action = "/employee/orders/order/next"
    } else if (value === "reject") {
        commentContainer.action = "/employee/orders/order/reject"
    }
}

nextBtn.addEventListener('click', () => {
    if (!nextBtn.classList.contains("inactive-btn")) {
        clickBtn("next")
    }
})

rejectBtn.addEventListener('click', () => {
    if (!rejectBtn.classList.contains("inactive-btn")) {
        clickBtn("reject")
    }
})