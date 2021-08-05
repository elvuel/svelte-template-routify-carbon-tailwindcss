export default function user(translate) {
    return {
        create: {
            name: {
                required: true
            },
            intro: {
                max: 30
            }
        }
    }
}