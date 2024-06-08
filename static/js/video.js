import { auth } from './auth.js';

var eContainer = document.querySelector('.container');
window.onload = async () => {
  if (await !auth) {
    alert('請先完成登錄');
    return;
  }

  fetch('/video/list', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json', // 设置内容类型为JSON
      Authorization: localStorage.getItem('access_token'),
    },
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.status == 'ok') {
        data.videoes.forEach((element) => {
          eContainer.insertAdjacentHTML(
            'beforeend',
            `
                <a href="/movie/${element.name}">
                    <div class="board">
                        <div class="left">
                            <img src="/static/img/${element.image}" alt="${element.image}">
                        </div>
                        <div class="right">
                            <div class="title"><p>${element.title}</p></div>
                            <div class="time"><p>${element.time}</p></div>
                            <div class="intro">
                                <p>${element.intro} ${element.name}</p>
                            </div>
                        </div>
                    </div>
                </a>
                `
          );
        });
      }
    })
    .catch((error) => {
      alert(error);
    });
};
