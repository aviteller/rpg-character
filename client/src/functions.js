export const getData = endpoint => {
  return fetch(`http://localhost:9000/api/${endpoint}`)
    .then(res => res.json())
    .then(data => data.data)
    .catch(err => console.log(err));
};

export const postData = (endpoint, data) => {
  return fetch(`http://localhost:9000/api/${endpoint}`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json" }
  })
    .then(res => res.json())
    .then(data => data.data)
    .catch(err => console.log(err));
};

// fetch("http://localhost:9000/api/class", {
//   method: "POST",
//   body: JSON.stringify(newClass),
//   headers: { "Content-Type": "application/json" }
// })
//   .then(res => {
//     if (!res.ok) {
//       throw new Error("Failed");
//     }
//     return res.json();
//   })
//   .then(data => {
//     if (!data.status) {
//       throw new Error(data.message);
//     }
//     title = "";
//     getClasses();
//   })
//   .catch(err => console.log(err));
