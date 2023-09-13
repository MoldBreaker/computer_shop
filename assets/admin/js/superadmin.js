
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
                html += `
                <tr>
                    <td>${users[i].user_name}</td>
                    <td>${users[i].address}</td>
                    <td>${users[i].phone}</td>
                    <td>${users[i].email}</td>
                    <td>${roles.map(each => {
                        if(each.role_id == users[i].role_id){
                            return each.role_name;
                        }
                    })}</td>
                    <td>${users[i].created_at}</td>
                    <td>${users[i].password == "" ? "Blocked" : "Active"}</td>
                    <td><button data-id="${users[i].user_id}" onclick="blockUser(this)" type="button" class="btn btn-danger">Block</button></td>
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

function blockUser(e){
    console.log(e.dataset.id)

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