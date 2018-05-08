FROM ruby

RUN mkdir -p /app
WORKDIR /app

COPY Gemfile .

RUN bundle install

COPY . .

EXPOSE 4567

CMD ruby helloworld.rb