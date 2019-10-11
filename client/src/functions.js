export const getClasses = () => {
  
  fetch(`http://localhost:9000/api/class`)
    .then(res => {
      if (!res.ok) {
        throw new Error("Issue fetching classes");
        
      }

      return res.json();
    })
    .then(data => {
      
      return data.data;
   
    })
    .catch(err => {
  
      console.log(err);
    });
};
