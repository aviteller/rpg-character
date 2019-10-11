<script>
  import { onMount } from "svelte";

  let loading = true;
  let classes = [];
  let title = "";

  const getClasses = () => {
    loading = true
    fetch(`http://localhost:9000/api/class`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching classes");
          loading = false;
        }
        return res.json();
      })
      .then(data => {
        if(data.data) {
        classes = data.data;
        }
        loading = false;
      })
      .catch(err => {
        loading = false;
        console.log(err);
      });
  };

  onMount(() => {
    getClasses();
  });

  const onSubmit = () => {
    const newClass = {
      title
    };

    fetch("http://localhost:9000/api/class", {
      method: "POST",
      body: JSON.stringify(newClass),
      headers: { "Content-Type": "application/json" }
    })
      .then(res => {
        if (!res.ok) {
          throw new Error("Failed");
        }
        return res.json();
      })
      .then(data => {
        if (!data.status) {
          throw new Error(data.message);
        }
        title = "";
        getClasses();
      })
      .catch(err => console.log(err));
  };
</script>

<div>
  <input type="text" bind:value={title} />
  <button on:click={onSubmit}>Add Class</button>
</div>

{#if classes.length > 0 && !loading}
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Title</th>
      </tr>
    </thead>
    <tbody>
      {#each classes as c}
        <tr>
          <td>{c.id}</td>
          <td>{c.title}</td>
        </tr>
      {/each}
    </tbody>
  </table>
{:else}
  <h1>Add Classes</h1>
{/if}
