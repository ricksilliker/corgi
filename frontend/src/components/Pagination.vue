<template>
<div class="container">
  <div class="info">
    Page {{ currentPage }}  of  {{ totalPages }}
  </div>
  <div class="controls">
    <div class="button">
      <feather type="arrow-left-circle" size="30px"></feather>
    </div>
    <div class="button">
      <feather type="arrow-right-circle" size="30px"></feather>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "Pagination",
  props: ["currentPage", "totalPages"]
}
</script>

<style scoped>
.container {
  width: 100%;
}

.controls {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
}

.button {
  color: white;
  width: 50px;
  height: 50px;
  cursor: pointer;
  padding: 5px;
}

.button:hover {
  color: darkgoldenrod;
}

.info {
  margin-top: 25px;
  text-align: center;
  font-size: 9pt;
  margin-bottom: 5px;
}
</style>