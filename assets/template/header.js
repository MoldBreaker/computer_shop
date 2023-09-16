let notificationBell = document.getElementById('notification-bell');

try{
    getAllNotifications();
} catch(e) {

}

try {
    renderCartCount();
}catch(e) {

}

try {
    notificationBell.onclick = () => {
        setNotifications();
    }
} catch (e) {

}

function getAllNotifications() {
    $.ajax({
        type: "GET",
        url: "/api/notifications/",
        dataType: "JSON",
        success: function (response) {
            if(response == null) {
                document.getElementById('notifications-count').classList.add('hidden');
            } else {
                document.getElementById('notifications-count').innerText = response.length
            }
        },
        error: function (jqXHR) {
            console.log(jqXHR)
            if(jqXHR.status === 403) {
                return false;
            }
        }
    });
}

function setNotifications() {
    $.ajax({
        type: "GET",
        url: "/api/notifications/",
        dataType: "JSON",
        success: function (response) {
            if(response == null) {
            } else {  
                let html = ''
                for(let i=0;i<response.length;i++){
                    html += `[<a style="color: red;" onclick="deleteNotification(this)" data-id="${response[i].NotificationId}" >DELETE</a>]<p>[${reformatDate(response[i].CreatedAt)}]: ${response[i].Content}</p><hr/>`
                }
                document.getElementById('notification-body').innerHTML = html
            }
        },
        error: function (jqXHR) {
            if(jqXHR.status === 400) {
                return jqXHR.responseJSON.message;
            }
        }
    });
}


function renderCartCount() {
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
}

function deleteNotification(e) {
    $.ajax({
        type: "DELETE",
        url: "api/notifications/" + e.dataset.id,
        dataType: "JSON",
        success: function (response) {
            setNotifications();
            getAllNotifications();
        }, 
        error: function (jqXHR) {
            console.log(jqXHR.status);
        }
    });
}

function reformatDate(dateString) {

    // Create a Date object from the input string
    const date = new Date(dateString);

    // Get the time components (hours, minutes, and seconds)
    const hours = date.getUTCHours().toString().padStart(2, '0');
    const minutes = date.getUTCMinutes().toString().padStart(2, '0');
    const seconds = date.getUTCSeconds().toString().padStart(2, '0');

    // Get the date components (month, day, and year)
    const month = (date.getUTCMonth() + 1).toString().padStart(2, '0'); // Add 1 to month because it's zero-based
    const day = date.getUTCDate().toString().padStart(2, '0');
    const year = date.getUTCFullYear();

    // Create the reformatted string
    return `${hours}:${minutes}:${seconds} ${day}/${month}/${year}`;
}