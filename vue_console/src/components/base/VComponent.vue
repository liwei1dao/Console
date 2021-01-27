<template>
  <section class="mb-12 text-center">
    <h1
      class="font-weight-light mb-2 headline"
      v-text="$t(heading)"
    />

    <span
      class="font-weight-light subtitle-1"
    >
     {{$t(destring)}}
    </span>
  </section>
</template>

<script>
  export default {
    name: 'VComponent',
    props: {
      heading: {
        type: String,
        default: '',
      },
      destring: {
        type: String,
        default: '',
      },
    },
  }
</script>
