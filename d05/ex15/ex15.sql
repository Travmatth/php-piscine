SELECT CONCAT(SUBSTR(`phone_number`,
         3),
         '5') AS 'rebmunenohp'
FROM `distrib`
WHERE `phone_number` LIKE '05%';