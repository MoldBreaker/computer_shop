/*!
* Start Bootstrap - Shop Homepage v5.0.6 (https://startbootstrap.com/template/shop-homepage)
* Copyright 2013-2023 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-shop-homepage/blob/master/LICENSE)
*/
// This file is intentionally blank
// Use this file to add JavaScript to your project
$(document).ready(function () {
    $.ajax({
        url: '/api/products/',
        type: 'GET',
        dataType: 'json',
        success: function (data, status, xhr) {
            console.log(data);
            html = ``
            for (let i = 0; i < data.length; i++) {
                html += `<div class="col mb-5">
                <div class="card h-100">
                        <!-- Product image-->
                        <img class="card-img-top"width="300px" src="${data[i].images==null?"https://dummyimage.com/450x300/dee2e6/6c757d.jpg":data[i].images[0]}" alt="${data[i].product_name}" />
                        <!-- Product details-->
                        <div class="card-body p-4">
                            <div class="text-center">
                                <!-- Product name-->
                                <h5 class="fw-bolder">${data[i].product_name}</h5>
                                <!-- Product price-->
                                ${formatMoney(data[i].price)}₫
                            </div>
                        </div>
                        <!-- Product actions-->
                        <div class="card-footer p-4 pt-0 border-top-0 bg-transparent">
                            <div class="text-center"><a class="btn btn-outline-dark mt-auto" href="#">Thêm vào giỏ hàng</a></div>
                        </div>
                    </div>
                </div>`
            }
            document.querySelector('#items-container').innerHTML = html
        },
        error: function (xhr, status, error) {
            console.log(xhr.status);
        }
    });
})


function formatMoney(number) {
    const absoluteNumber = Math.abs(number);
    const absNumberWithCommas = absoluteNumber.toLocaleString('en-US');

    const parts = absNumberWithCommas.split('.');
    const wholePart = parts[0];
    const decimalPart = parts[1] || '';

    const formattedWholePart = wholePart.replace(/\B(?=(\d{3})+(?!\d))/g, '.');

    if (decimalPart === '') {
        return (number < 0 ? '-' : '') + formattedWholePart;
    } else {
        return (number < 0 ? '-' : '') + formattedWholePart + '.' + decimalPart;
    }
}