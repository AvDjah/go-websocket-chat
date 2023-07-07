<template>
  <div class="font-mono">
    <button class="bg-green-300 p-4 ring-2 ring-zinc-400 rounded-lg m-2" @click="props.getPools">Get All Pools</button>
    <div class="border-2 m-2 p-2">
      <br /> <span class="text-lg mb-2 p-2">Pool list</span>
      <div>
        <li class="pool-id-item cursor-pointer ring-red-200 ring-4 rounded-xl bg-blue-300 p-6 pb-8 m-4  md:w-1/4"
          v-for="(pool, index) of pool_list" @click="(event) => props.sendPoolId(event, pool)">
          > {{ pool.name }}
        </li>
      </div>
    </div>
  </div>
</template>


<script setup>

import gsap from 'gsap'


const pool_list = useState("pool_list")

const props = defineProps({
  getPools: Function,
  sendPoolId: Function
})

// const send_pool_id = (event, id) => {
//   console.log(id.pool_id)
//   const resp = {
//     type: "POOL_ID_JOIN",
//     data: id.pool_id
//   }
//   webSocket.send(JSON.stringify(resp))
// }


</script>