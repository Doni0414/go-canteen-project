var createBtn = document.getElementById('create-btn')
var exitBtns = document.getElementsByClassName('exit-btn')

Array.from(exitBtns).forEach((btn, index) => {
    btn.addEventListener('click', () => {
        document.getElementById('darken').classList.remove('active')
        document.getElementsByClassName('product-action-container')[index].classList.remove('active')
    })
})

createBtn.addEventListener('click', () => {
    document.getElementById('darken').classList.add('active')
    document.getElementById('create-product-container').classList.add('active')
})

var editBtns = document.getElementsByClassName('edit-btn')
var productIdContainers = document.getElementsByClassName('product-id-container')
var xhr = new XMLHttpRequest()

Array.from(editBtns).forEach((btn, index) => {
    btn.addEventListener('click', () => {
        document.getElementById('darken').classList.add('active')
        document.getElementById('edit-product-container').classList.add('active')

        xhr.open('GET', 'http://localhost:8080/employee/products?id=' + productIdContainers[index].value, true)

        xhr.onreadystatechange = function() {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    var responseData = JSON.parse(xhr.responseText)['product']
                    document.getElementById('edit-product-id').value = responseData['ID']
                    document.getElementById('edit-product-name').value = responseData['Name']
                    document.getElementById('edit-product-category-id').value = responseData['CategoryId']
                    document.getElementById('edit-product-price').value = responseData['Price']
                }
            }
        }

        xhr.send()
    })
})

var createProductPriceInput = document.getElementById('create-product-price')
createProductPriceInput.addEventListener('input', () => {
    createProductPriceInput.value = Math.abs(createProductPriceInput.value)
})

$('#create-product-btn').click(function() {
    var formData = new FormData()
    formData.append("name", $('#create-product-name').val())
    formData.append("category_id", $('#create-product-category-id').val())
    formData.append('price', $('#create-product-price').val())
    formData.append('image', $('#create-product-image')[0].files[0])

    $.ajax({
        url: "/employee/products/create",
        type: "POST",
        data: formData,
        processData: false,
        contentType: false,
        success: function(response) {
            console.log("Form submitted successfully");
            window.location.replace(`http://localhost:8080/employee/main?q=${document.getElementById('search-q').value}&category=${document.getElementById('search-category-id').value}&page=${document.getElementById('search-page').value}`)
        },
        error: function(xhr, status, error) {
            console.error("Error:", error);
            var response = xhr.responseJSON
            var errorMessage = document.getElementById('creation-error-message')
            errorMessage.style.display = 'block'
            errorMessage.textContent = response['product_creation_error']
        }
    })
})

$('#edit-product-btn').click(function() {
    var formData = new FormData()
    formData.append('id', $('#edit-product-id').val())
    formData.append('name', $('#edit-product-name').val())
    formData.append('category_id', $('#edit-product-category-id').val())
    formData.append('price', $('#edit-product-price').val())
    formData.append('image', $('#edit-product-image')[0].files[0])

    $.ajax({
        url: "/employee/products/edit",
        type: "POST",
        data: formData,
        processData: false,
        contentType: false,
        success: function(response) {
            console.log("Form submitted successfully");
            window.location.replace(`http://localhost:8080/employee/main?q=${document.getElementById('search-q').value}&category=${document.getElementById('search-category-id').value}&page=${document.getElementById('search-page').value}`)
        },
        error: function(xhr, status, error) {
            console.error("Error:", error);
            var response = xhr.responseJSON
            var errorMessage = document.getElementById('edit-error-message')
            errorMessage.style.display = 'block'
            errorMessage.textContent = response['product_edit_error']
        }
    })
})