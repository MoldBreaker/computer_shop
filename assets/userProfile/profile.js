let logoutBtn = document.getElementById('logout-btn')
let changeAvatarbtn = document.getElementById('change-avatar-save-btn')
let updateInfoBtn = document.getElementById('update-information-save-btn')
let validImageExt = ["jpeg", "png", "jpg", "jfif"]
let regexPhoneNumber = /(84|0[3|5|7|8|9])+([0-9]{8})\b/g;
let changePasswordbtn = document.getElementById('change-password-save-btn')
const isNonWhiteSpace = /^\S*$/;
const isContainsUppercase = /^(?=.*[A-Z]).*$/;
const isContainsLowercase = /^(?=.*[a-z]).*$/;
const isContainsNumber = /^(?=.*[0-9]).*$/;
const isContainsSymbol = /^(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]).*$/;
const isValidLength = /^.{6,16}$/;
let regularExpression = /^(\S)(?=.*[0-9])(?=.*[A-Z])(?=.*[a-z])(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹])[a-zA-Z0-9~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]{6,16}$/;
var currentUrl = window.location.href;
var match = currentUrl.match(/\/(\d+)\/?$/);
let currentUserId = document.getElementById('current-id').value;

$.ajax({
  type: "GET",
  url: "/api/users/" + match[1],
  dataType: "JSON",
  success: function (response) {
    let user = response.user;
    let role = response.role;
    const dateString = user.created_at;
    const dateArray = dateString.split("T"); 
    const datePortion = dateArray[0];
    const parts = datePortion.split("-");
    const formattedDate = `${parts[2]}/${parts[1]}/${parts[0]}`;
    let html = `
    <div
      class="ms-4 mt-5 d-flex flex-column cover object circular-image"
      style="width: 150px"
    >
      ${user.avatar == '' ? `<img
      src="/home/img/user.png"
      alt="Generic placeholder image"
      class="img-fluid img-thumbnail mt-4 mb-2"
      style="z-index: 1; border-radius: 50%"
    />` : `<img
      src="${user.avatar}"
      alt="Generic placeholder image"
      class="img-fluid img-thumbnail mt-4 mb-2 contain"
      style="
        z-index: 1;
        border-radius: 50%;
        width: 150px;
        height: 150px;
      "
    />`}
    <div
      style="
        z-index: 1;
        display: flex;
        justify-content: space-between;
      "
    >
      ${currentUserId == user.user_id ? `<button
      type="button"
      class="btn btn-outline-dark"
      data-mdb-ripple-color="dark"
      style="z-index: 1; margin-left: 16px"
      data-bs-toggle="modal"
      data-bs-target="#form-avatar"
    >
      Change Avatar
    </button>
    <button
      type="button"
      class="btn btn-outline-dark"
      data-mdb-ripple-color="dark"
      style="z-index: 1; margin-left: 16px"
      data-bs-toggle="modal"
      data-bs-target="#update-infomation-form"
    >
      Update Information
    </button>
    <button
      type="button"
      class="btn btn-outline-dark"
      data-mdb-ripple-color="dark"
      style="z-index: 1; margin-left: 16px"
      data-bs-toggle="modal"
      data-bs-target="#changePassword"
    >
      Change Password
    </button>
    <button
      type="button"
      class="btn btn-outline-dark"
      data-mdb-ripple-color="red"
      style="z-index: 1; margin-left: 16px"
      id="logout-btn" onclick="handlerLogout()"
    >
      Logout
    </button>` : ``}
    </div>
  </div>
    <div class="ms-3" style="margin-top: 130px">
      <h3>${user.user_name} (${role.role_name})</h3>
      Joined At:
      <a id="created-at">${formattedDate}</a>
    </div>
    `;
    document.getElementById('userinfo').innerHTML = html;
  },
  error: function (jqXHR) {
    if(jqXHR.status === 400){
      document.getElementById('info-container').innerHTML = `<h1 style="text-align: center; color: red;">${jqXHR.responseJSON.message}</h1>`;
    }
  }
});

function handlerLogout() {
  Swal.fire({
    title: 'Are you sure?',
    text: "Do you want to logout?",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Yes'
  }).then((result) => {
    if (result.isConfirmed) {
      $.ajax({
        type: "GET",
        url: "/api/users/logout",
        success: function (data) {
            window.location.href = "/"
        },
        error: function (jqXHR) {
            console.log(jqXHR.responseText);
        }
      });
    }
  })
}

async function handlerChangeAvatar() {
  let check = true;
  let changeAvatarForm = document.getElementById('change-avatar-form');
  formAvatar = new FormData(changeAvatarForm);
  file = formAvatar.get('avatar');
  let fileExt = file.name.split('.').pop();
  if(file.name.length === 0){
    document.getElementById("avatar-error").innerHTML = "Please select a file";
    check = false;
  } else if(!validImageExt.includes(fileExt)){
    document.getElementById("avatar-error").innerHTML = "Please must be jpeg, png, jpg, jfif extension";
    check = false;
  } else {
    document.getElementById("avatar-error").innerHTML = '';
  }
  if(!check) {
    e.preventDefault();
    return false;
  } else {
    await changeAvatarForm.submit();
  }
}

function handlerUpdateInfo() {
  let check = true;
  formHtml = document.getElementById('info-form');
  let form = new FormData(formHtml);

  let phone = form.get('phone');
  let address = form.get('address');

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

  if(!check) return false;

  $.ajax({
    type: "POST",
    url: "/api/users/info",
    data: {
      phone: phone,
      address: address
    },
    dataType: "JSON",
    success: function (response) {
      Swal.fire(
        'Nice!',
        'Your information have been updated!',
        'success'
      )
    },
    error: function (jqXHR){
      Swal.fire({
        icon: 'error',
        title: 'Something when wrong',
        text: `${jqXHR.responseJSON.message}`,
      })
    }
  });
  return;
}

function handlerChangePassword() {
  let check = true;
  let changePasswordHtml = document.getElementById('changepass-form');
  let formData = new FormData(changePasswordHtml);
  let oldPassword = formData.get('oldpassword');
  let newPassword = formData.get('newpassword');
  let confirmNewPassword = formData.get('cnewpassword');


  if(!isNonWhiteSpace.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must not contain Whitespaces.'
    check = false;
  }
  if (!isContainsUppercase.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must have at least one Uppercase Character.'
    check = false;
  }
  if (!isContainsLowercase.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must have at least one Lowercase Character.'
    check = false;
  }
  if(!isContainsNumber.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must have at least one Number'
    check = false;
  }
  if(!isContainsSymbol.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must have at least one Special Symbol.'
    check = false;
  }
  if(!isValidLength.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='Password must be 6-16 Characters Long.'
    check = false;
  }
  if(regularExpression.test(newPassword)){
    document.getElementById('new-password-error').innerHTML ='';
  }


  if(confirmNewPassword != newPassword){
    document.getElementById('confirm-new-password-error').innerHTML = 'Passwords do not match.'
    check = false;
  }else{
      document.getElementById('confirm-new-password-error').innerHTML = ''
  }

  if(check==false){
    return false;
  }

  $.ajax({
    type : "POST",
    url : "/api/users/reset-password",
    data : {
      old_password : oldPassword,
      new_password : newPassword,
      confirm_new_password : confirmNewPassword
    },
    dataType: "JSON",
    success: function (response) {
      Swal.fire(
          'Nice!',
          'Password changed successfully!',
          'success'
      )
    },
    error: function (changePassword){
      document.getElementById('old-password-error').innerHTML = `Please enter the correct password`
    }
  })
  return;
}