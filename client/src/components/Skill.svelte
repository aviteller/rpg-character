<script>
  import { onMount } from "svelte";
  import Select from "./Select.svelte";

  let loading = true;
  let skills = [];
  let attributes = [];
  let name = "";
  let attribute_id = "";

  const getAttributes = () => {
    fetch(`http://localhost:9000/api/attribute`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching attributes");
        }

        return res.json();
      })
      .then(data => {
        if (data.data) {
          attributes = data.data;
          attributes.forEach(c => {
            c["value"] = c["id"];
            c["label"] = c["name"];
            delete c["id"];
            delete c["name"];
          });
        }
      })
      .catch(err => {
        console.log(err);
      });
  };

  const getSkills = () => {
    loading = true;
    fetch(`http://localhost:9000/api/skills`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching skills");
          loading = false;
        }
        return res.json();
      })
      .then(data => {
        if (data.data) {
          skills = data.data;
        }
        loading = false;
      })
      .catch(err => {
        loading = false;
        console.log(err);
      });
  };

  onMount(async () => {
    await getAttributes();
    await getSkills();
  });

  const handleSelect = (e, field) => {
    switch (field) {
      case "attribute":
        attribute_id = e.detail;
        break;
      case "class":
        class_id = e.detail;
        break;
    }
  };

  const onSubmit = () => {
    const newSkill = {
      name,
      attribute_id
    };

    console.log(newSkill);

    fetch("http://localhost:9000/api/skills", {
      method: "POST",
      body: JSON.stringify(newSkill),
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
        modifier = 0;
        getTraits();
      })
      .catch(err => console.log(err));
  };

  const getClassName = class_id => {
    if (classes.length > 0) {
      let className = classes.filter(c => c.value == class_id)[0].label;
      return className;
    }
  };

  const getAttributeName = attribute_id => {
    if (attributes.length > 0) {
      let attributeName = attributes.filter(a => a.value == attribute_id)[0]
        .label;
      return attributeName;
    }
  };
</script>

<style>
  .form {
    display: flex;
    justify-content: space-between;
    max-width: 50%;
  }
</style>

{#if !loading}
  <div class="form">
    <Select
      options={attributes}
      on:selectchange={e => handleSelect(e, 'attribute')} />
    <input type="text" bind:value={name} />
    <div>

      <button on:click={onSubmit}>Add Trait</button>
    </div>
  </div>
{/if}

{#if skills.length > 0 && !loading}
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>name</th>
      </tr>
    </thead>
    <tbody>
      {#each skills as c}
        <tr>
          <td>{c.id}</td>
          <td>{getAttributeName(c.attribute_id)}</td>
          <td>{c.name}</td>
        </tr>
      {/each}
    </tbody>
  </table>
{:else}
  <h1>Add skills</h1>
{/if}
