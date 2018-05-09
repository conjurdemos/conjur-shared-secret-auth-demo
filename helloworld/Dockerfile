FROM ruby:2.5

RUN mkdir -p /app
WORKDIR /app

COPY Gemfile .
RUN bundle install

COPY . .

EXPOSE 4567

CMD ruby helloworld.rb