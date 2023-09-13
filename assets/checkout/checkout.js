let checkoutBtn = document.getElementById('checkout-btn')
let regexPhoneNumber = /(84|0[3|5|7|8|9])+([0-9]{8})\b/g;
let cartArray = [];

checkoutBtn.onclick = (e) =>{
    e.preventDefault();
    let check = true;
    let formHtml = document.getElementById('checkout-form')
    let formData = new FormData(formHtml)
    let phone = formData.get('phone')
    let address = formData.get('address')

    if(phone.length == 0) {
        document.getElementById('phone-error').innerHTML = 'Please enter your phone number';
        check = false;
    } else if(!phone.match(regexPhoneNumber)) {
        document.getElementById('phone-error').innerHTML = 'Invalid phone number';
        check = false;
    } else {
        document.getElementById('phone-error').innerHTML = '';
    }

    if(address.length == 0) {
        document.getElementById('address-error').innerHTML = 'Please enter your address';
        check = false;
    } else {
        document.getElementById('address-error').innerHTML = '';
    }

    console.log(JSON.stringify({
        cart: cartArray
    }))

    if(!check) return false;

    Swal.fire({
        title: 'Have you checked the information is correct?',
        showDenyButton: true,
        confirmButtonText: 'Yes',
        denyButtonText: `No`,
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                method: 'POST',
                url: '/api/users/info',
                data: {
                    phone: phone,
                    address: address
                },
                dataType: 'json',
                success: function(response) {
                    $.ajax({
                        method: 'POST',
                        url: '/api/invoices/',
                        contentType: 'application/json',
                        dataType: 'json',
                        data: JSON.stringify({
                            cart: cartArray
                        }),
                        success: function(data) {
                            window.location.href = '/'
                        },
                        error: function(error) {
                            Swal.fire(
                                'Nice!',
                                `${error.responseJSON.message}`,
                                'error'
                            )
                        }
                    })
                },
                error: function(jqXHR){
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.responseText,
                    })
                }
            })
        }
    })


}
$.ajax({
    method: 'GET',
    url : '/api/carts/',
    dataType: 'JSON',
    success: function(data) {
        let html = ''
        html += `<h4>Cart <span class="price" style="color:black"><i class="fa fa-shopping-cart"></i> <b>${data.length}</b></span></h4>`
        let sum = 0;
        for(let i = 0; i < data.length; i++){
            html += `<p><a href="/product/detail/${data[i].Product.product_id}">${data[i].Product.product_name}</a> x${data[i].Quantity} <span class="price">${formatMoney(data[i].Price)}</span></p>`;
            cartArray.push([data[i].Product.product_id, data[i].Quantity])
            sum+=data[i].Price
        }
        html += `<p>Total <span class="price" style="color:black"><b>${formatMoney(sum)}</b></span></p>`
        document.getElementById('container').innerHTML = html;

    },

    error: function(jqXHR){
        console.log(jqXHR);
    }

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