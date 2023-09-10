let logoutBtn = document.getElementById('logout-btn')
let changeAvatarbtn = document.getElementById('change-avatar-save-btn')
let updateInfoBtn = document.getElementById('update-information-save-btn')
let validImageExt = ["jpeg", "png", "jpg", "jfif"]
let regexPhoneNumber = /(84|0[3|5|7|8|9])+([0-9]{8})\b/g;
const dateString = document.getElementById('created-at').innerText;
const dateArray = dateString.split(" "); 
const datePortion = dateArray[0];
const parts = datePortion.split("-"); 


const formattedDate = `${parts[2]}/${parts[1]}/${parts[0]}`;
document.getElementById('created-at').innerText = formattedDate

logoutBtn.onclick = () => {
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

changeAvatarbtn.onclick = async (e) => {
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

updateInfoBtn.onclick = (e) => {
  e.preventDefault();
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