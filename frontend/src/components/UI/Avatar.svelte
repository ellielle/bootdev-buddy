<script>
  export let image;
  export let handle;
</script>

<main>
  <label for="img-profile">
    <img src={image} alt={handle} id="img-profile" />
  </label>
</main>

<style>
  #img-profile {
    width: 40px;
    height: auto;
    border-radius: 50%;
  }
</style>
