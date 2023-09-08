/*!
* Start Bootstrap - Shop Homepage v5.0.6 (https://startbootstrap.com/template/shop-homepage)
* Copyright 2013-2023 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-shop-homepage/blob/master/LICENSE)
*/
// This file is intentionally blank
// Use this file to add JavaScript to your project
let url = new URL(window.location.href);
let params = new URLSearchParams(url.search);
let page = params.get('page') || 1;
let productsLength
let searchBtn = document.querySelector("#search-btn")
let sortBtnList = document.querySelectorAll(".sort-li")
let col = ""
let sort = ""

$(document).ready(function () {

    for(let i=0;i<sortBtnList.length;i++){
        sortBtnList[i].onclick = (e) =>{
            e.preventDefault();
            col = e.target.dataset.col;
            sort = e.target.dataset.sort;
            params.set('col', col);
            params.set('sort', sort);
            renderListProducts();
        }
    }
    renderListProducts();
    searchBtn.onclick = function () {
        let searchValue = document.getElementById("search-value").value;
        if(searchValue.length == 0){
            document.getElementById("search-warning").innerHTML = 'Please enter something'
            return;
        } else {
            document.getElementById("search-warning").innerHTML = ''
        }
        params.set('search', searchValue);
        renderListProducts();
    }
})

const renderListProducts = () => {
    $.ajax({
        url: '/api/products/?' + params.toString(),
        type: 'GET',
        dataType: 'json',
        success: function (data, status, xhr) {
            productsLength = data.maxLength;
            data = data.products;
            html = ``
            for (let i = 0; i < data.length; i++) {
                html += `<div class="col mb-5">
                <div class="card h-100">
                        <!-- Product image-->
                        <img class="card-img-top hover-jumpin" width="300px" src="${data[i].images==null?"https://dummyimage.com/450x300/dee2e6/6c757d.jpg":data[i].images[0]}" alt="${data[i].product_name}" />
                       
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
            $('#page').Pagination({ // id to initial draw and use pagination
                size: productsLength, // size of list input
                pageShow: 5, // 5 page-item per page
                page: page, // current page (default)
                limit: 8, // current limit show perpage (default)
            }, function(obj){ // callback function, you can use it to re-draw table or something
                params.set('page', obj.page)
                renderListProducts();
                //window.location.href = url.origin + '?' + params.toString()
                window.scrollTo(0, 0);
            });
        },
        error: function (err) {
            if(err.status == 400){
                html = `<h1>${err.responseJSON.message}</h1>`
                document.querySelector('#items-container').innerHTML = html
            }
        }
    });
}


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

document.addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
        event.preventDefault();
        searchBtn.click();
    }
});



function myFunction() {
    document.getElementById("myDropdown").classList.toggle("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function(event) {
    if (!event.target.matches('.dropbtn')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}


