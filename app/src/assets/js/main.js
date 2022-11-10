const displayList = function(e) {
  e.preventDefault();

  const data = new FormData();
  const text = document.getElementById("register-text").value;
  data.append("text", text);

  axios
    .post("/api", data)
    .then((res) => {
      const res_list = res.data;
      document.querySelectorAll(".memo-list")[0].innerHTML = rewriteList(res_list);
    })
    .catch((error) => {
      console.log(error);
    });
};

const button = document.getElementById("register-button");

window.addEventListener("load", displayList);
button.addEventListener("click", displayList);

function rewriteList($res_list) {
  let $list = '<ul class="item-list">\n';
  for( var $i = 0; $i < $res_list.length; $i++ ) {
    $list += '<li class="list-item">\n';
    $list += `<div class="memo-date">${$res_list[$i].date}</div>\n`
    $list += `<div class="memo-content">${$res_list[$i].text}</div>\n`
    $list += '</li>\n';
  }
  $list += '</ul>';
  return $list;
}