var slideIndex = 1;
var url = window.location.href;
var match = url.match(/\/(\d+)\/?$/);

$.ajax({
  type: "GET",
  url: "/api/carts/",
  dataType: "JSON",
  success: function (data) {
      document.getElementById('cart-count').innerHTML = data.length;
  },
  error: function (jqXHR){
      document.getElementById('cart-count').innerHTML = 0;
  }
});

$.ajax({
  type: "GET",
  url: "/api/products/" + match[1],
  dataType: "JSON",
  success: function (data) {
    let imagesHTML = "";
    for(let i=0;i<data.images.length;i++){
      imagesHTML += `
      <img
        class="mySlides"
        src="${data.images[i]}"  ${i==0? "style='display: block;'" : ""}
      />
      `
    }
    imagesHTML += `<button
        class="w3-button w3-black w3-display-left"
        onclick="plusDivs(-1)"
      >
        &#10094;
      </button>
      <button
        class="w3-button w3-black w3-display-right"
        onclick="plusDivs(1)"
      >
        &#10095;
      </button>`;
    document.getElementById("slide-content").innerHTML = imagesHTML;

    let html  = `
    <div class="product-detail-card">
          <div class="product-name">
            <h1 class="title">
              ${data.product_name}
            </h1>
          </div>
          <div class="product-price">Price: ${formatMoney(data.price)}₫</div>
          <div class="add-to-cart-btn">
            <button data-id="${data.product_id}" onclick="addTocart(this)" type="button" class="btn btn-primary">Add to Cart</button>
          </div>
          <div class="product-description">
            <h3 class="description-title">Description:</h3>
            <p class="description">
              ${data.description}
            </p>
          </div>
          <div class="back-btn">
            <a href="/" class="btn btn-primary">Back to Home</a>
          </div>
        </div>
    `
    document.getElementById("right-item").innerHTML = html;
  }
});

showDivs(slideIndex);

function plusDivs(n) {
  showDivs(slideIndex += n);
}

function addTocart(e) {
  $.ajax({
      type: "GET",
      url: "/api/carts/" + e.dataset.id,
      dataType: "JSON",
      success: function (response) {
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
          renderCartCount();
          Swal.fire(
          'Nice!',
          'Add to Cart successfully',
          'success'
          )
      },
      error: function (jqXHR) {
          if (jqXHR.status === 403) {
              window.location.href = "/auth"
          }
      }
  });
}

function showDivs(n) {
  var i;
  var x = document.getElementsByClassName("mySlides");
  if (n > x.length) {slideIndex = 1}
  if (n < 1) {slideIndex = x.length}
  for (i = 0; i < x.length; i++) {
    x[i].style.display = "none";  
  }
  x[slideIndex-1].style.display = "block";  
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
