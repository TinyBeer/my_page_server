import { auth, login } from "./auth.js"

var modal = document.getElementById("login_modal");
var btn_ok = document.getElementById("ok");
var form = document.querySelector('form');
var btn_close = document.getElementById("close");
var out = document.getElementById("log_out");
var input_username = document.getElementById("username");
var input_password = document.getElementById("password");
var isLogin = false; // 标记是否正在登录

let menuToggle = document.querySelector('.menuToggle');
let menu = document.querySelector('.menu');


async function nav_login() {
    if (isLogin) return; // 如果已经在登录，则直接返回
    var username = input_username.value;
    var password = input_password.value;
    var res = await login(username, password);
    if (res.ok) {
        modal.style.display = "none";
        isLogin = true;
        menu.classList.toggle('active');
    } else {
        alert(res.mes);
    }

    input_username.value = "";
    input_password.value = "";
    return;
}

input_username.onkeydown = (event) => {
    if (event.key == 'Enter') {
        input_password.focus();
    }
    return;
}

input_password.onkeydown = (event) => {
    if (event.key == 'Enter') {
        nav_login();
    }
    return;
}

form.onsubmit = function (e) {
    e.preventDefault(); // 阻止表单的默认提交行为
};

btn_ok.onclick = nav_login;
btn_close.onclick = function () {
    input_username.value = "";
    input_password.value = "";
    modal.style.display = "none"; // 关闭登录窗口
};

out.onclick = function () {
    var userResponse = window.confirm("你确定要注销吗？");
    if (userResponse) {
        console.log("用户点击了确定。");
        localStorage.removeItem('access_token');
        localStorage.removeItem('refresh_token');
        isLogin = false;
        menu.classList.toggle('active');
    } else {
        console.log("用户点击了取消。");
    }
};

menuToggle.onclick = function () {
    // menu.classList.toggle('active');
    if (isLogin) return;
    modal.style.display = "block";
    input_username.focus();
};

window.onload = async function () {
    console.log("页面加载完毕");
    // 页面加载完成验证token
    // 如果有可用的token直接进入
    if (isLogin) {
        menu.classList.toggle('active');
    }

    if (await auth()) {
        console.log("验证token成功");
        isLogin = true;
        menu.classList.toggle('active');
        return
    }
    console.log("验证token失败");
    // 没有可用的token 需要点击十字进行登录
};
