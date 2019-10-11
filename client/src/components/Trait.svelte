<script>
  import { onMount } from "svelte";
  import Select from "./Select.svelte";

  let loading = true;
  let traits = [];
  let classes = [];
  let attributes = [];
  let modifier = 0;
  let class_id = "";
  let attribute_id = "";

  const getClasses = () => {
    fetch(`http://localhost:9000/api/class`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching classes");
        }
        return res.json();
      })
      .then(data => {
        if (data.data) {
          classes = data.data;
          classes.forEach(c => {
            c["value"] = c["id"];
            c["label"] = c["title"];
            delete c["id"];
            delete c["title"];
          });
        }
      })
      .catch(err => {
        console.log(err);
      });
  };

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

  const getTraits = () => {
    loading = true;
    fetch(`http://localhost:9000/api/traits`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching traits");
          loading = false;
        }
        return res.json();
      })
      .then(data => {
        if (data.data) {
          traits = data.data;
        }
        loading = false;
      })
      .catch(err => {
        loading = false;
        console.log(err);
      });
  };

  onMount(async () => {
    await getClasses();
    await getAttributes();
    await getTraits();
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
    const newTrait = {
      modifier,
      class_id,
      attribute_id
    };

    console.log(newTrait);

    fetch("http://localhost:9000/api/traits", {
      method: "POST",
      body: JSON.stringify(newTrait),
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
      let attributeName = attributes.filter(a => a.value == attribute_id)[0].label;
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
    <Select options={classes} on:selectchange={e => handleSelect(e, 'class')} />
    <Select
      options={attributes}
      on:selectchange={e => handleSelect(e, 'attribute')} />
    <input type="number" bind:value={modifier} />
    <button on:click={onSubmit}>Add Class</button>
  </div>
{/if}

{#if traits.length > 0 && !loading}
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>modifier</th>
      </tr>
    </thead>
    <tbody>
      {#each traits as c}
        <tr>
          <td>{c.id}</td>
          <td>{getClassName(c.class_id)}</td>
          <td>{getAttributeName(c.attribute_id)}</td>
          <td>{c.modifier}</td>
        </tr>
      {/each}
    </tbody>
  </table>
{:else}
  <h1>Add traits</h1>
{/if}
