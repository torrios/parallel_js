
var app = new Vue({
    el: "#app",
    data: {
        name: "Hector Rios"
    },
    methods: {
        startTest() {
            var $main = $('#main');
            $('#button').click(function () {
                var number = $('#number').val(),
                delay = $('#delay').val();
                for (var i = 0; i < number; i++) {
                    var start = (new Date()).getTime();
                    // $.ajax({
                    //     url: '<?php echo basename($_SERVER['PHP_SELF']); ?>',
                    //     data: {'ajax': 1, 'delay': delay},
                    //     async: true,
                    //     success: function (data) {
                    //         var end = (new Date()).getTime(),
                    //         elapsed = end - start;
                    //         $main.append('<div>' + data + '  - Elapsed time is: ' + elapsed + ' ms</div>');
                    //     }
                    // });
                }
            });
        },

        clearResults() {
            $main.html('');
        }
    }
})
