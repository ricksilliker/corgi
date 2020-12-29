<template>
<div class="container">
  <div class="rightSide">
    <div class="actionButton">
      <feather class="actionButtonIcon" type="settings"></feather>
    </div>
    <div class="actionButton">
      <feather class="actionButtonIcon" type="bell"></feather>
    </div>
  </div>
</div>
</template>

<script>
export default {
name: "Header"
}
</script>

<style scoped>
.container {
  border-bottom: 1px solid white;
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;
  height: 80px;
}

.rightSide {
  box-sizing: border-box;
  position: absolute;
  top: 0;
  right: 0;
  width: 10%;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5px;
}

.actionButton {
  flex: 1 1 30px;
  color: white;
  cursor: pointer;
  position: relative;
  margin: 10px;
}

.actionButton:hover {
  color: darkgoldenrod;
}

.actionButtonIcon {
  height: 30px;
  width: 30px;
}
</style>