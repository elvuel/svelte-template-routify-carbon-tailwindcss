export default function user(translate) {
    return {
        create: {
            name: {
                required: true
            },
            intro: {
                max: 30
            }
        },
        update: {
            name: {
                required: true
            },
            intro: {
                max: 30
            }
        }
    }
}