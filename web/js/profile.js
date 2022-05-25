const app = {
    data() {
        return {
            show_main: false,
            show_input: true,
            add_password: false,
            show_about: false,
            inputPassword: null,
            documentation: false,
            aboutself: false,
            inputPasswordConfirm: null,
            file: null
        }
    },
    methods: {
        dark_checked() {
            const check_box = document.querySelector('.dark-theme');

            check_box.addEventListener('click', function(e) {
                if (this.checked) {
                    document.body.classList.add('dark-theme');
                } else {
                    document.body.classList.remove('dark_theme');
                }
            })
        },
        submit_button_file() {
            const check_box = document.querySelector('.dark-theme');
        }

    }
}

Vue.createApp(app).mount('#app-main')