var userArray = [];
function AJAX(
  { url, type },
  JsonData,
  callback,
  headers = { "Content-Type": "application/json", Authorization: "Bearer" }
) {
  jQuery.ajax({
    url,
    type,
    dataType: "json",
    data: JSON.stringify(JsonData),
    headers,
    success: function (data) {
      callback(data, null);
    },
    error: function (err) {
      callback(null, err);
    },
  });
}
jQuery(document).ready(function () {
  // BEGIN login button event
  jQuery("#submit").click(function () {
    let mUserName = jQuery("#fname").val();
    let mEmail = jQuery("#email").val();
    let mPassword = jQuery("#password").val();
    AJAX(
      { url: "http://localhost:8080/api/v1/login", type: "post" },
      {
        Username: mUserName,
        Password: mPassword,
        Email: mEmail,
      },
      (data, err) => {
        if (err) {
          console.log(err);
          return;
        }
        // This is token.
        sessionStorage.setItem("test2@3", data.message.Token);
        sessionStorage.setItem("testUuid", data.message.Uuid);
        console.log(data);
        jQuery("#dispUuid").html(data.message.Uuid);
        alert(data.message.Uuid + "\n" + data.message.Token);
        jQuery("#signup").show();
        return;
      }
    );
  });
  //[] END login

  // BEGIN Signup button event
  jQuery("#signup").click(function () {
    let mUserName = jQuery("#fname").val();
    let mEmail = jQuery("#email").val();
    let mPassword = jQuery("#password").val();
    AJAX(
      { url: "http://localhost:8080/api/v1/register", type: "post" },
      {
        Username: mUserName,
        Password: mPassword,
        Email: mEmail,
        Uuid: sessionStorage.getItem("testUuid"),
      },
      (data, err) => {
        if (err) {
          console.log(err);
          alert("User ID is duplicated.");
          return;
        }
        alert(data.message.Uuid + " is signed up");
      }
    );
  });
  //[] END login

  // BEGIN update by user
  jQuery("#users").click(function () {
    let mUserName = jQuery("#fname").val();
    let mEmail = jQuery("#email").val();
    let mPassword = jQuery("#password").val();
    AJAX(
      { url: "http://localhost:8080/api/v1/users", type: "post" },
      {
        Username: mUserName,
        Password: mPassword,
        Email: mEmail,
        Uuid: sessionStorage.getItem("testUuid"),
      },
      (data, err) => {
        if (err) {
          console.log(err);
          alert("Error");
          return;
        }
        alert("good");
      },
      {
        "Content-Type": "application/json",
        Authorization: "Bearer " + sessionStorage.getItem("test2@3"),
      }
    );
  });
  //[] END update by user

  // BEGIN get all user info
  jQuery("#getAll").click(function () {
    AJAX(
      {
        url: "http://localhost:8080/api/v1/users",
        type: "get",
      },
      null,
      (data, err) => {
        if (err) {
          console.log(err);
          alert("No exist data");
          return;
        }
        console.log(data);
        userArray = data.message;
        alert("Torch [Draw Table] button, pleas");
      },
      {
        "Content-Type": "application/json",
        Authorization: "Bearer " + sessionStorage.getItem("test2@3"),
      }
    );
  });
  //[] get all
  // BEGIN get user info
  jQuery("#getUser").click(function () {
    let mUuid = jQuery("#userID").val();
    if (mUuid === "") {
      alert("Input user id and try again");
      return;
    }
    AJAX(
      {
        url: "http://localhost:8080/api/v1/users?user_id=" + mUuid,
        type: "get",
      },
      null,
      (data, err) => {
        if (err) {
          console.log(err);
          alert("No exist data");
          return;
        }
        console.log(data);
        userArray = data.message || [];
        alert("Torch [Draw Table] button, pleas");
      },
      {
        "Content-Type": "application/json",
        Authorization: "Bearer " + sessionStorage.getItem("test2@3"),
      }
    );
  });
  //[] get user
  // BEGIN delete user info
  jQuery("#delUser").click(function () {
    let mUuid = jQuery("#userID").val();
    if (mUuid === "") {
      alert("Input user id and try again");
      return;
    }
    AJAX(
      {
        url: "http://localhost:8080/api/v1/users?user_id=" + mUuid,
        type: "delete",
      },
      {},
      (data, err) => {
        if (err) {
          console.log(err);
          alert("User ID is duplicated.");
          return;
        }
        console.log(data);
        alert(data.message);
      },
      {
        "Content-Type": "application/json",
        Authorization: "Bearer " + sessionStorage.getItem("test2@3"),
      }
    );
  });
  //[] delete
  // BEGIN update user info
  jQuery("#updateUser").click(function () {
    let mUuid = jQuery("#userID").val();
    let mUsername = jQuery("#u_username").val();
    let mEmail = jQuery("#u_email").val();
    let mApiUID = jQuery("#u_api_user_id").val();
    if (mUuid === "" || mUsername === "" || mEmail === "" || mApiUID === "") {
      alert("Input all data and try again");
      return;
    }
    AJAX(
      {
        url: "http://localhost:8080/api/v1/users?user_id=" + mUuid,
        type: "post",
      },
      { Username: mUsername, Email: mEmail, ApiUserID: mApiUID },
      (data, err) => {
        if (err) {
          console.log(err);
          alert("User ID is duplicated.");
          return;
        }
        console.log(data);
        alert(data.message);
      },
      {
        "Content-Type": "application/json",
        Authorization: "Bearer " + sessionStorage.getItem("test2@3"),
      }
    );
  });
  //[] delete

  // Draw Table
  jQuery("#drawTable").click(function () {
    var html = "";
    userArray.map((item) => {
      html += `<tr>
        <td>${item.Uuid}</td><td>${item.UserName}</td><td>${item.Email}</td><td>${item.ApiUserID}</td>
      </tr>`;
    });
    jQuery("#tbody4user").html(html);
  });
});
