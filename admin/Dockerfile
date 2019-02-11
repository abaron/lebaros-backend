FROM php:7.2-fpm

# set working directory
WORKDIR /var/www

# install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    mysql-client \
    libpng-dev \
    libjpeg62-turbo-dev \
    libfreetype6-dev \
    locales \
    zip \
    jpegoptim optipng pngquant gifsicle \
    vim \
    nano \
    unzip \
    git \
    curl

# clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# install extensions
RUN docker-php-ext-install pdo_mysql mbstring zip exif pcntl
RUN pecl install -o -f redis \
    &&  rm -rf /tmp/pear \
    &&  docker-php-ext-enable redis
RUN apt-get update && apt-get install -y \
    libmagickwand-dev --no-install-recommends \
    && pecl install imagick \
    && docker-php-ext-enable imagick

# install composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# add user for yii2 application
RUN groupadd -g 1000 www
RUN useradd -u 1000 -ms /bin/bash -g www www

# copy existing application directory contents
RUN mkdir admin
COPY . /var/www/admin

# set working directory
WORKDIR /var/www/admin

# install dependency
RUN pwd
RUN ls -la
RUN composer install

# copy existing application directory permissions
COPY --chown=www:www . /var/www

# change current user to www
USER www

# expose port 9000 and start php-fpm server
EXPOSE 9000
CMD ["php-fpm"]