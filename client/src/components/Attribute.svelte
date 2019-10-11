<script>
  import { onMount } from "svelte";

  let loading = true;
  let attributes = [];
  let name = "";

  const getAttributes = () => {
    loading = true;
    fetch(`http://localhost:9000/api/attribute`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching attributes");
          loading = false;
        }

        return res.json();
      })
      .then(data => {
        if (data.data) {
          attributes = data.data;
        }
        loading = false;
      })
      .catch(err => {
        loading = false;
        console.log(err);
      });
  };

  onMount(() => {
    getAttributes();
  });

  const onSubmit = () => {
    const newClass = {
      name
    };

    //console.log(newClass)
    fetch("http://localhost:9000/api/attribute", {
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
        name = "";
        getAttributes();
      })
      .catch(err => console.log(err));
  };
</script>

<div>
  <input type="text" bind:value={name} />
  <button on:click={onSubmit}>Add Attribute</button>
</div>

{#if attributes.length > 0 && !loading}
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Title</th>
      </tr>
    </thead>
    <tbody>
      {#each attributes as c}
        <tr>
          <td>{c.id}</td>
          <td>{c.name}</td>
        </tr>
      {/each}
    </tbody>
  </table>
{:else}
  <h1>Add attributes</h1>
{/if}
