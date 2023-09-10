

$(document).ready(function() {
    renderCart();
})

function renderCart() {
    $.ajax({
        type: "GET",
        url: "/api/carts/",
        dataType: "JSON",
        success: function (data) {
            let html = ''
            let totalPrice = 0;
            for(let i=0;i<data.length;i++) {
                totalPrice += data[i].Price
                html += `
                <div class="card rounded-3 mb-4">
                <div class="card-body p-4">
                  <div
                        class="row d-flex justify-content-between align-items-center"
                    >
                        <div class="col-md-2 col-lg-2 col-xl-2">
                        <img
                            src="${data[i].Product.images[0]}"
                            class="img-fluid rounded-3"
                            alt="${data[i].Product.product_name}"
                        />
                        </div>
                        <div class="col-md-3 col-lg-3 col-xl-3">
                        <a href="/product/detail/${data[i].Product.product_id}">
                            <p class="lead fw-normal mb-2">${data[i].Product.product_name}</p>
                        </a>
                        </div>
                        <div class="col-md-3 col-lg-3 col-xl-2 d-flex">
                        <button
                            class="btn btn-link px-2"
                            onclick="decrease(this)"
                        >
                            <i class="fas fa-minus"></i>
                        </button>
    
                        <input
                            id="form1"
                            min="0"
                            name="quantity"
                            data-id="${data[i].Product.product_id}"
                            value="${data[i].Quantity}"
                            type="number"
                            class="form-control form-control-sm"
                            disabled
                        />
    
                        <button
                            class="btn btn-link px-2"
                            onclick="increase(this)"
                        >
                            <i class="fas fa-plus"></i>
                        </button>
                        </div>
                        <div class="col-md-3 col-lg-2 col-xl-2 offset-lg-1">
                        <h5 class="mb-0">${formatMoney(data[i].Price)}</h5>
                        </div>
                        <div class="col-md-1 col-lg-1 col-xl-1 text-end">
                        <a data-id="${data[i].Product.product_id}" onclick="removeInCart(this)" class="text-danger"
                            ><i class="fas fa-trash fa-lg"></i
                        ></a>
                        </div>
                    </div>
                    </div>
                </div>
                `
            }
            document.getElementById('cart-item-container').innerHTML = html;
            document.getElementById('total-price').innerHTML = formatMoney(totalPrice);

            $.ajax({
                type: "GET",
                url: "/api/carts/",
                dataType: "JSON",
                success: function (data) {
                    document.getElementById('cart-count').innerHTML = data.length;
                },
                error: function (jqXHR){
                    console.log(jqXHR.responseJSON);
                }
            });
        }, 
        error: function (jqXHR) {
            if(jqXHR.status === 400){
                document.getElementById('cart-item-container').innerHTML = `<h1>${jqXHR.responseJSON.message} <a href="/">Buy one now</a></h1>`;
                document.getElementById('checkout-btn').style.display = 'none';
                document.getElementById('total-price').innerHTML = 0;
            }
        }
    });
}

function increase(e) {
    let input = e.parentNode.querySelector('input[type=number]');
    let id = input.dataset.id;
    $.ajax({
        type: "GET",
        url: "/api/carts/update/" + id + "?type=increase",
        dataType: "JSON",
        success: function (response) {
            e.parentNode.querySelector('input[type=number]').stepUp();
            renderCart();
        },
        error: function (jqXHR){
            console.log(jqXHR.responseText);
        }
    });
}

function decrease(e) {
    let input = e.parentNode.querySelector('input[type=number]');
    let id = input.dataset.id;
    $.ajax({
        type: "GET",
        url: "/api/carts/update/" + id + "?type=decrease",
        dataType: "JSON",
        success: function (response) {
            e.parentNode.querySelector('input[type=number]').stepDown()
            renderCart();
        },
        error: function (jqXHR){
            console.log(jqXHR.responseText);
        }
    });
}

function removeInCart(e) {
    let id = e.dataset.id;
    $.ajax({
        type: "DELETE",
        url: "/api/carts/" + id,
        dataType: "JSON",
        success: function (response) {
            renderCart();
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