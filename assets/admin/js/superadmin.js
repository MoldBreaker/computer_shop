
$(document).ready(function() {
    renderUserTable();
})

function renderUserTable() {
    $.ajax({
        method : 'GET',
        url : '/api/users/',
        dataType : 'json',
        success : function(data) {
            let html = ''
            let users = data.users
            let roles = data.roles
            for (let i = 0; i < users.length;i++){
                const dateString = users[i].created_at;
                const dateArray = dateString.split("T"); 
                const datePortion = dateArray[0];
                const parts = datePortion.split("-");
                const formattedDate = `${dateArray[1].split("Z")[0]} ${parts[2]}/${parts[1]}/${parts[0]}`;
                html += `
                <tr>
                    <td><a style="text-decoration: none;" href="/profile/${users[i].user_id}">${users[i].user_name}</a></td>
                    <td>${users[i].address}</td>
                    <td>${users[i].phone}</td>
                    <td>${users[i].email}</td>
                    <td>${roles.map(each => {
                        if(each.role_id == users[i].role_id){
                            return each.role_name;
                        }
                    })}</td>
                    <td>${formattedDate}</td>
                    <td>${users[i].password == "" ? "Blocked" : "Active"}</td>
                    <td>
                        <button data-id="${users[i].user_id}" onclick="blockUser(this)" type="button" class="btn btn-danger">Block</button>
                        <button data-id="${users[i].user_id}" data-role="${users[i].role_id}" onclick="updateRole(this)" type="button" class="btn btn-primary">Update Role</button>
                    </td>
                </tr>
                `
            }
            document.querySelector("tbody").innerHTML = html;
        },
        error: function(jqXHR) {
            console.log(jqXHR.responseJSON.message);
        }
    })
}
function getRolesObject(e) {
    $.ajax({
        type: "GET",
        url: "/api/role/",
        dataType: "JSON",
        success: function (response) {
            let data = {};
            for(let i=0; i<response.roles.length; i++){
                if(response.roles[i].role_name == 'super_admin' || response.roles[i].role_id == e.dataset.role){
                    continue;
                } else {
                    data[response.roles[i].role_id] = response.roles[i].role_name;
                }
            }
            selectOptionSwal(data, e.dataset.id)
        }
    });
}

function selectOptionSwal(data, userId) {
    Swal.fire({
        title: 'Select Role',
        input: 'select',
        inputOptions: data,
        inputPlaceholder: 'required',
        showCancelButton: true,
        inputValidator: function (value) {
          return new Promise(function (resolve, reject) {
            if (value !== '') {
              resolve();
            } else {
              resolve('You need to select a Role');
            }
          });
        }
      }).then(function (result) {
        if (result.isConfirmed) {
          $.ajax({
            type: "PUT",
            url: "/api/role/" + userId,
            data: {
                role_id: result.value
            },
            dataType: "JSON",
            success: function (response) {
                window.location.reload();
            }, 
            error: function (jqXHR){
                if(err.status == 400){
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: jqXHR.responseJSON.message,
                    })
                }
            }
          });
        }
      });
}

function updateRole(e) {
    getRolesObject(e)
}

function blockUser(e){

    Swal.fire({
        title: 'Do you want to block this user, that can not be undone?',
        showDenyButton: true,
        confirmButtonText: 'Yes',
        denyButtonText: `No`,
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                method: "PUT",
                url: "/api/users/block/" + e.dataset.id,
                dataType: "json",
                success: function(data){
                    window.location.reload();
                },
                error: function(jqXHR) {
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.responseJSON.message,
                    })
                }
            })
        }
    })
}