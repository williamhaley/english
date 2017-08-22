(() => {

$('form').on('submit', function () {
	const $form = $(this);

	const url = $form.data('url');

	$(this).trigger('reset');

	return false;
});

})()