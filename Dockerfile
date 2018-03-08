FROM amelia/php:fpm

COPY . /srv/code

RUN composer install --no-dev \
    && php artisan route:cache \
    && php artisan view:clear \
    && php artisan storage:link \
    && rm -rf /var/cache/composer/*