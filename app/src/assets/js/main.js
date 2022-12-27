const button = document.getElementById("register-button");

const displayList = function (e) {
  e.preventDefault();

  const data = new FormData();
  const text = document.getElementById("register-text").value;
  console.log(text);
  data.append("text", text);

  axios
    .post("/api", data)
    .then((res) => {
      button.disabled = true;
      showModal(text);
      const res_list = res.data;
      if (null !== res_list) {
        document.querySelectorAll(".memo-list")[0].innerHTML = rewriteList(res_list);
      }

      button.disabled = false;
    })
    .catch((error) => {
      console.log(error);
    });
};

window.addEventListener("load", displayList);
button.addEventListener("click", displayList);

const closeButton = document.querySelector(".modal-close");
closeButton.addEventListener("click", function () {
  document.querySelector(".modal").style.display = "none";
});

function rewriteList($res_list) {
  let $list = '<ul class="item-list">\n';
  for (var $i = 0; $i < $res_list.length; $i++) {
    $list += '<li class="list-item">\n';
    $list += `<div class="memo-date">${$res_list[$i].date}</div>\n`;
    $list += `<div class="memo-content">${$res_list[$i].text}</div>\n`;
    $list += "</li>\n";
  }
  $list += "</ul>";
  return $list;
}

function showModal(text) {
  if (text) {
    const addMessage = `アイデア「${text}」を鍋に入れました！`;
    document.querySelector(".modal-text").innerText = addMessage;
    document.querySelector(".modal").style.display = "block";
  }
}
