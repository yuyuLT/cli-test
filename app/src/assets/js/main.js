const button = document.getElementById("register-button");

button.addEventListener("click", (e) => {
  e.preventDefault();

  const data = new FormData();
  const text = document.getElementById("register-text").value;
  data.append("text", text);

  axios
    .post("/api", data)
    .then((res) => {
      console.log(res);
      const memo_list = res.data;
    })
    .catch((error) => {
      console.log(error);
    });
});
