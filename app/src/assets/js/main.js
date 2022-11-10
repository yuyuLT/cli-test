const button = document.getElementById("register-button");

button.addEventListener("click", (e) => {
  e.preventDefault();

  const data = new FormData();
  const text = document.getElementById("register-text").value;
  data.append("text", text);

  axios
    .post("/api", data)
    .then((res) => {
      const res_list = res.data;
      Object.keys(res_list).forEach(function (key) {
        console.log('date : ' + res_list[key].date);
        console.log('text : ' +res_list[key].text);
      });

    })
    .catch((error) => {
      console.log(error);
    });
});
