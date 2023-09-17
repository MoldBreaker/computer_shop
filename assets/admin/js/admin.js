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
let categoryId = ""
let addProductBtn = document.querySelector('#add-product-btn')
let modifyBtn = document.querySelector('#modify-btn')
let productForm = document.querySelector('#product-form')
let validImageExt = ["jpeg", "png", "jpg", "jfif"]
let closeBtn = document.querySelector('#close-btn')

$(document).ready(function () {

    $.ajax({
        type: "GET",
        url: "/api/categories/",
        dataType: "JSON",
        success: function (data) {
            let html = ''
            for(let i=0;i<data.categories.length;i++){
                html += `<a class="category-li" onclick="findByCategoryId(this)" data-id="${data.categories[i].CategoryId}">${data.categories[i].CategoryName}</a>`
            }
            document.getElementById("categories-dropdown").innerHTML = html;
        }
    });

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

    addProductBtn.onclick = () => {
        setCategoriesIntoForm("");
    }
    modifyBtn.onclick = () => {
        handleSubmit();
    }
    closeBtn.onclick = () => {
        handleClose();
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
                html += `<div
                class="col mt-4"
                style="display: flex; justify-content: center; align-items: center"
              >
                <div class="card" style="width: 18rem">
                  <img src="${data[i].images==null?"https://dummyimage.com/450x300/dee2e6/6c757d.jpg":data[i].images[0]}" alt="${data[i].product_name}" class="card-img-top" />
                  <div class="card-body">
                    <h5 class="card-title">
                    <a class="product-detail-link" href="/product/detail/${data[i].product_id}">
                        ${data[i].product_name}
                    </a>
                    </h5>
                    <p class="card-text">
                        ${formatMoney(data[i].price)}₫
                    </p>
                    <a data-id="${data[i].product_id}" onclick="handleUpdate(this)" data-bs-toggle="modal" data-bs-target="#modify-modal" class="btn btn-primary">Edit</a>
                    <a data-id="${data[i].product_id}" data-name="${data[i].product_name}" onclick="handleDelete(this)" class="btn btn-danger">Delete</a>
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

function handleClose() {
    document.getElementById('inputProductName').value = "";
    document.getElementById('inputProductDescription').innerText = "";
    document.getElementById('inputProductPrice').value = "";
    setCategoriesIntoForm()
    document.getElementById('images').disabled = false;
    document.getElementById('modify-btn').innerText = "Create";
}

function setCategoriesIntoForm(categoryId) {
    $.ajax({
        type: "GET",
        url: "/api/categories/",
        dataType: "JSON",
        success: function (data) {
            let html = ''
            for(let i=0;i<data.categories.length;i++){
                if(data.categories[i].CategoryId == categoryId){
                    html += `<option selected value="${data.categories[i].CategoryId}">${data.categories[i].CategoryName}</option>`
                } else {
                    html += `<option value="${data.categories[i].CategoryId}">${data.categories[i].CategoryName}</option>`
                }
            }
            document.getElementById("category-list").innerHTML = html;
        }
    });
}

function handleUpdate(e) {
    $.ajax({
        type: "GET",
        url: "/api/products/" + e.dataset.id,
        dataType: "JSON",
        success: function (response) {
            document.getElementById('inputProductName').value = response.product_name;
            document.getElementById('inputProductDescription').innerText = response.description;
            document.getElementById('inputProductPrice').value = response.price;
            setCategoriesIntoForm(response.category_id)
            document.getElementById('images').disabled = true;
            document.getElementById('modify-btn').innerText = "Modify";
            document.getElementById('product-id').value = e.dataset.id;
        }
    });
}

function handleDelete(e) {
    Swal.fire({
        title: 'Are you sure?',
        text: `Are you sure you want to delete ${e.dataset.name} from products?`,
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes'
      }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                type: "DELETE",
                url: "/api/products/" + e.dataset.id,
                dataType: "JSON",
                success: function (response) {
                    window.location.reload();
                }, 
                error: function (xhr) {
                    if(xhr.status != 200) {
                        Swal.fire(
                            'Error',
                            xhr.responseText,
                            'error'
                          )
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
      })
}

function handleSubmit() {
    let check = true;
    let formData = new FormData(productForm);
    let productId = formData.get('product_id');
    let productName = formData.get('product_name');
    let categoryId = formData.get('category_id');
    let description = formData.get('description');
    let price = formData.get('price');
    let imagesInput = document.getElementById('images');
    let images = imagesInput.files;

    if(productName.length == 0){
        document.getElementById('name-error').innerText = "Please enter a name";
        check = false; 
    } else {
        document.getElementById('name-error').innerText = '';
    }

    if(description.length == 0){
        document.getElementById('description-error').innerText = "Please enter a description";
        check = false; 
    } else {
        document.getElementById('description-error').innerText = '';    
    }

    if(categoryId == "") {
        document.getElementById('category-error').innerText = "Please select a category";
        check = false;
    }  else {
        document.getElementById('category-error').innerText = "";
    }

    if(price.length == 0) {
        document.getElementById('price-error').innerText = "Please enter a price";
        check = false;
    } else if(!/^\d*\.?\d+$/.test(price)) {
        document.getElementById('price-error').innerText = "Invalid price";
        check = false;
    } else {
        document.getElementById('price-error').innerText = "";
    }

    if(document.getElementById('images').disabled == false){
        if(images.length == 0) {
            document.getElementById('images-error').innerText = "Please select at least one image";
            check = false;
        } else {
            let checkImage = true;
            for(let i=0;i<images.length;i++) {
                let fileExt = images[i].name.split('.').pop();
                if(!validImageExt.includes(fileExt)){
                    document.getElementById("images-error").innerHTML = "Images must be jpeg, png, jpg, jfif extension";
                    checkImage = false;
                }
            }
            if(!checkImage){
                check = false;
            }
        }
    } 

    if(!check){
        e.preventDefault();
        return false;
    }
    if(modifyBtn.innerText == "Create"){
        productForm.submit();
    } else {
        $.ajax({
            type: "PUT",
            url: "/api/products/" + productId,
            data: {
                product_id: productId,
                category_id: categoryId,
                product_name: productName,
                description: description,
                price: price
            },
            dataType: "JSON",
            success: function (response) {
                window.location.reload();
            },
            error: function (jqXHR){
                if(jqXHR.status != 200) {
                    Swal.fire(
                        'Error',
                        jqXHR.responseJSON.message,
                        'error'
                      )
                }
            }
        });
    }
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

function showCategories() {
    document.getElementById("categories-dropdown").classList.toggle("show");
}

function findByCategoryId(e) {
    categoryId = e.dataset.id;
    params.set("categoryId", categoryId);
    renderListProducts();
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


