<template>
	<transition name="message-fade">
		<div class="message" :class="typeClass" role="alert" v-show="visible">
			<p class="message__content">
				<v-icon large :color="color" :small="true">{{icon}}</v-icon>  &nbsp;{{message}}
			</p>
		</div>
	</transition>
</template>

<script>
	export default {
		name: 'message',
		props : {
			visible: {
				type: Boolean,
				default:false,
			},
			duration: {
				type: Number,
				default:2000,
			},
			message: {
				type: String,
				default:"",
			},
			timer: {
				type: Object,
				default:null,
			},
			closed: {
				type: Boolean,
				default:false,
			},
			color: {
				type: String,
				default:"",
			},
			icon: {
				type: String,
				default:"",
			},
			typeClass: {
				type: String,
				default:"",
			},
		},

		watch: {
			closed(val) {
				if (val) {
					this.visible = false
					this.$el.addEventListener('transitionend', this.destroyElement)
				}
			}
		},

		methods: {
			destroyElement() {
				this.$el.removeEventListener('transitionend', this.destroyElement)
				this.$destroy(true)
				this.$el.parentNode.removeChild(this.$el)
			},

			startTimer() {
				if (this.duration > 0) {
					this.timer = setTimeout(() => {
						if (!this.closed) {
							this.close()
						}
					}, this.duration)
				}
			},

			close() {
				this.closed = true
			}
		},

		mounted() {
			// 开始定时器
			this.startTimer()
		}
	}
</script>

<style scoped>
	.message {
		min-width: 380px;
		box-sizing: border-box;
		border-radius: 4px;
		border-width: 1px;
		border-style: solid;
		/* border-color: #ebeef5; */
		position: fixed;
		left: 50%;
		top: 20px;
		transform: translateX(-50%);
		/* background-color: #edf2fc; */
		transition: opacity .3s, transform .4s;
		overflow: hidden;
		padding: 15px 15px 15px 20px;
		display: flex;
		align-items: center
	}

	.message--success {
		background-color: #f0f9eb;
		border-color: #e1f3d8;
		color: #67c23a;
	}

    .message--alert {
        /*背景颜色*/
        background-color: #f0f9eb;
        /*边框颜色*/
        border-color: #e1f3d8;
        /*文字颜色*/
        color: #67c23a;
    }
	
	.message--error {
	background-color: #FFCCFF;
	border-color: #FFCCFF;
	color: #CC0033;
	}

	.message p {
		margin: 0
	}

	.message__content {
		padding: 0;
		font-size: 14px;
		line-height: 1;
	}

	.message-fade-enter,
	.message-fade-leave-active {
		opacity: 0;
		transform: translate(-50%, -100%)
	}
</style>