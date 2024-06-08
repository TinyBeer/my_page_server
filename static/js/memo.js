import { auth } from './auth.js';

var eload = document.querySelector('.load_wrapper');
var ebtn = document.querySelector('#btn_add');
var ebox = document.querySelector('.box');
var einp = document.querySelector('#note_content');

window.onload = async function () {
  if (!(await auth())) {
    console.log('login first');
  }
  let intervalId = setInterval(function () {
    auth();
  }, 1000 * 60 * 3);
  await refreshMemoList();
  eload.classList.add('fade-out');
};

ebtn.onclick = async function () {
  if (await createMemo(einp.value)) {
    refreshMemoList();
  }
  einp.value = '';
  einp.focus();
};
ebox.onclick = async function (event) {
  var tg = event.target;
  if (tg.innerHTML == 'x') {
    var id = parseInt(tg.parentElement.parentElement.firstElementChild.value);
    if (await deleteMemo(id)) {
      refreshMemoList();
    }
  }
};

async function refreshMemoList() {
  await listMemo()
    .then((data) => {
      if (data.status == 'ok') {
        while (ebox.firstChild) {
          ebox.removeChild(ebox.firstChild);
        }

        if (data.memoes) {
          data.memoes.forEach((element) => {
            ebox.insertAdjacentHTML(
              'beforeend',
              `<div class="note">
                        <input type="hidden" value="${element.id}">
                <a href="#">
                    <span>x</span>
                    <p>${element.content}</p>
                </a>
            </div>`
            );
          });
        }
      }
    })
    .catch((error) => alert(error));
}

async function listMemo() {
  var access_token = localStorage.getItem('access_token');
  return await fetch('/memo/list', {
    method: 'GET', // 使用GEt方法发送数据
    headers: {
      // 'Content-Type': 'application/json', // 设置内容类型为JSON
      Authorization: access_token,
    },
  })
    .then((response) => response.json())
    .catch((error) => console.error('list memo error:', error));
}

async function deleteMemo(id) {
  var isOk = false;
  var access_token = localStorage.getItem('access_token');
  await fetch('/memo/delete', {
    method: 'DELETE', // 使用GEt方法发送数据
    headers: {
      'Content-Type': 'application/json', // 设置内容类型为JSON
      Authorization: access_token,
    },
    body: JSON.stringify({
      id,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.status == 'ok') {
        isOk = true;
      } else {
        console.log(data);
      }
    })
    .catch((error) => console.error('delete memo error:', error));
  return isOk;
}

async function updateMemo(id, content) {
  var isOk = false;
  var access_token = localStorage.getItem('access_token');
  await fetch('/memo/update', {
    method: 'PUT', // 使用GEt方法发送数据
    headers: {
      // 'Content-Type': 'application/json', // 设置内容类型为JSON
      Authorization: access_token,
    },
    body: JSON.stringify({
      id,
      content,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.status == 'ok') {
        isOk = true;
      }
    })
    .catch((error) => console.error('update memo error:', error));
  return isOk;
}

async function createMemo(content) {
  var isOk = false;
  var access_token = localStorage.getItem('access_token');
  await fetch('/memo/create', {
    method: 'POST', // 使用POST方法发送数据
    headers: {
      'Content-Type': 'application/json', // 设置内容类型为JSON
      Authorization: access_token,
    },
    body: JSON.stringify({
      content,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.status == 'ok') {
        isOk = true;
      } else {
        console.log(data);
      }
    })
    .catch((error) => console.log('create memo error:', error));
  return isOk;
}
