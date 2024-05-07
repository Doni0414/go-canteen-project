var productQuantities = document.getElementsByClassName('product-quantity')
var productTotalPrices = document.getElementsByClassName('product-total-price')
var totalPriceContainer = document.getElementById('total-price')
var cartIdContainers = document.getElementsByClassName('cart-id-container')
var xhr = new XMLHttpRequest();

Array.from(productQuantities).forEach((el, index) => {
    el.addEventListener('input', event => {
        el.value = Math.abs(el.value)

        xhr.open('POST', 'http://localhost:8080/user/cart/changeQuantity', true);
        
        var data = {
            "cart-id": parseInt(cartIdContainers[index].value),
            "quantity": parseInt(el.value)
        }
        xhr.onreadystatechange = function () {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    var responseData = JSON.parse(xhr.responseText)
                    productTotalPrices[index].textContent = responseData["product-total-price"] + " ₸"
                    totalPriceContainer.textContent = responseData["total-price"] + " ₸"
                } else {
                    console.error('Request failed with status:', xhr.status);
                }
            }
        };

        xhr.send(JSON.stringify(data));
    })
});