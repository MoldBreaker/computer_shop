let logoutBtn = document.getElementById('logout-btn')
let changeAvatarbtn = document.getElementById('change-avatar-save-btn')
let validImageExt = ["jpeg", "png", "jpg", "jfif"]

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