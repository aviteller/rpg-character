<script>
  import { onMount } from "svelte";
  import Class from "./components/Class.svelte";
  import Attribute from "./components/Attribute.svelte";
  import Trait from "./components/Trait.svelte";
  import Skill from "./components/Skill.svelte";
  import { getData } from "./functions";

  let classes = [];
  let classesLoaded = false;
  let skills = [];
  let attributes = [];
  let traits = [];

  onMount(() => {
    getData("classes").then(res => {
			classes = res;
			//console.log(classes)
      classesLoaded = true;
    });
    getData("skills").then(res => (skills = res));
    getData("attributes").then(res => (attributes = res));
    getData("traits").then(res => (traits = res));
  });
</script>

<style>

</style>

{#if classesLoaded}
  <Class classes={classes} />
{/if}
<!--  <Attribute/>
<Trait/>
<Skill/> -->
