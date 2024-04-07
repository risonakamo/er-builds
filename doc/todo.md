# use cases
- trying to make a build, and want to see for each item category, what is the list of item other people have used, and how many people have used them
    - would like to select to filter by a certain version
    - for each item, sort by:
        - number of builds that has it
        - total number of likes
        - average win rate?
        - highest win rate?
- when clicking on an item, want to see all builds that include that item
    - this can show if this item is used with certain other items
    - sort by:
        - number of likes
        - win rate
        - date
- the following will need collection of user match data, since it doesn't seem to appear in the build lists
    - want to be able to see all the augments people are picking
    - want to see the tac skills people are picking
    - want to see yellow items being picked
    - maybe the process to collect data from user matches kind of overrides the build list?

# idea
- what if, we select various ppl's builds and look at their games to collect additional information like yellow iteams, augments and tac skill
    - program ideally go to every build in the list and collect the game history data for each person. but this would be pretty api expensive, even if doing it only once in a while. so what's the best way to do this?
        - pretty sure dak does it by collecting once in a while from the main game api, then holding onto the data. so pretty much i would have to also do the same in order to not rely on constantly grabbing the data

# todo
- [ ] tally up likes for each item
- [ ] data save - not really needed to continuously be pinging the api. can store data and only refresh from api once in a while
    - ability to write data and load it
- [ ] total percentage of results - for each item, currently counting the total. would like to know the percentage of the results this item takes up

# current issues
- [ ] average win rate doesn't really make sense, as a huge number of random builds with 0 win rate by default pull it down
    - think about better way to display win rate
- [x] for each item, need to parse it's item category from the description