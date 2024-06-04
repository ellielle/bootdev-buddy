<script>
  export let image;
  export let handle;
</script>

<main>
  <label for="img-profile">
    <a href="https://www.boot.dev/u/{handle}" target="_blank">
      <img src={image} alt={handle} id="img-profile" />
    </a>
  </label>
</main>

<style>
  #img-profile {
    width: 40px;
    height: auto;
    border-radius: 50%;
  }
</style>
